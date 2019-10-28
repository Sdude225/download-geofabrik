package main

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/generator"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version = "devel"

var (
	app      = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.")
	fService = app.Flag("service",
		"Can switch to another service. "+
			"You can use \"geofabrik\", \"openstreetmap.fr\" or \"bbbike\". "+
			"It automatically change config file if -c is unused.").
		Default("geofabrik").String()
	fConfig     = app.Flag("config", "Set Config file.").Default("./geofabrik.yml").Short('c').String()
	fNodownload = app.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	fVerbose    = app.Flag("verbose", "Be verbose").Short('v').Bool()
	fQuiet      = app.Flag("quiet", "Be quiet").Short('q').Bool()
	fProgress   = app.Flag("progress", "Add a progress bar").Bool()

	list = app.Command("list", "Show elements available")
	lmd  = list.Flag("markdown", "generate list in Markdown format").Bool()

	d          = app.Command("download", "Download element") // TODO : add d as command
	delement   = d.Arg("element", "OSM element").Required().String()
	dosmBz2    = d.Flag(formats.FormatOsmBz2, "Download osm.bz2 if available").Short('B').Bool()
	dosmGz     = d.Flag(formats.FormatOsmGz, "Download osm.gz if available").Short('G').Bool()
	dshpZip    = d.Flag(formats.FormatShpZip, "Download shp.zip if available").Short('S').Bool()
	dosmPbf    = d.Flag(formats.FormatOsmPbf, "Download osm.pbf (default)").Short('P').Bool()
	doshPbf    = d.Flag(formats.FormatOshPbf, "Download osh.pbf").Short('H').Bool()
	dstate     = d.Flag(formats.FormatState, "Download state.txt file").Short('s').Bool()
	dpoly      = d.Flag(formats.FormatPoly, "Download poly file").Short('p').Bool()
	dkml       = d.Flag(formats.FormatKml, "Download kml file").Short('k').Bool()
	dCheck     = d.Flag("check", "Control with checksum (default) Use --no-check to discard control").Default("true").Bool()
	dOutputDir = d.Flag("output_directory", "Set output directory, you can use also OUTPUT_DIR env variable").Short('d').String()

	generate = app.Command("generate", "Generate a new config file")
)

func listAllRegions(c *config.Config, format string) {
	table := tablewriter.NewWriter(os.Stdout)
	keys := make(sort.StringSlice, len(c.Elements))
	i := 0

	for k := range c.Elements {
		keys[i] = k
		i++
	}

	keys.Sort()
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ShortName", "Is in", "Long Name", "formats"})

	if format == "Markdown" {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}

	for _, item := range keys {
		table.Append([]string{item,
			c.Elements[c.Elements[item].Parent].Name,
			c.Elements[item].Name,
			formats.MiniFormats(c.Elements[item].Formats)})
	}

	table.Render()
	fmt.Printf("Total elements: %#v\n", len(c.Elements))
}

func checkService() bool {
	switch *fService {
	case "geofabrik":
		return true
	case "openstreetmap.fr":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./openstreetmap.fr.yml"
		}

		return true
	case "bbbike":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./bbbike.yml"
		}

		return true
	}

	return false
}

func listCommand() {
	var format = ""

	configPtr, err := config.LoadConfig(*fConfig)
	if err != nil {
		log.Panicln(err)
	}

	if *lmd {
		format = "Markdown"
	}

	listAllRegions(configPtr, format)
}

func downloadCommand() {
	formatFile := formats.GetFormats()

	configPtr, err := config.LoadConfig(*fConfig)
	if err != nil {
		log.Panicln(err)
	}

	for _, format := range *formatFile {
		if ok, _, _ := config.IsHashable(configPtr, format); *dCheck && ok {
			if fileExist(*dOutputDir + *delement + "." + format) {
				if !downloadChecksum(format) {
					if !*fQuiet {
						log.Println("Checksum mismatch, re-downloading", *dOutputDir+*delement+"."+format)
					}

					myElem, err := config.FindElem(configPtr, *delement)
					if err != nil {
						log.Panicln(err)
					}

					myURL, err := config.Elem2URL(configPtr, myElem, format)
					if err != nil {
						log.Panicln(err)
					}

					err = download.FromURL(myURL, *dOutputDir+*delement+"."+format)
					if err != nil {
						log.Panicln(err)
					}

					downloadChecksum(format)
				} else if !*fQuiet {
					log.Printf("Checksum match, no download!")
				}
			} else {
				myElem, err := config.FindElem(configPtr, *delement)
				if err != nil {
					log.Panicln(err)
				}
				myURL, err := config.Elem2URL(configPtr, myElem, format)
				if err != nil {
					log.Panicln(err)
				}
				err = download.FromURL(myURL, *dOutputDir+*delement+"."+format)
				if err != nil {
					log.Panicln(err)
				}
				if !downloadChecksum(format) && !*fQuiet {
					log.Println("Checksum mismatch, please re-download", *dOutputDir+*delement+"."+format)
				}
			}
		} else {
			myElem, err := config.FindElem(configPtr, *delement)
			if err != nil {
				log.Panicln(err)
			}
			myURL, err := config.Elem2URL(configPtr, myElem, format)
			if err != nil {
				log.Panicln(err)
			}
			err = download.FromURL(myURL, *dOutputDir+*delement+"."+format)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
}

func configureBool(flag *bool, name string) {
	viper.Set(name, false)

	if *flag {
		viper.Set(name, true)
	}
}

func main() {
	app.Version(version) // Add version flag
	commands := kingpin.MustParse(app.Parse(os.Args[1:]))

	configureBool(fVerbose, "verbose")
	configureBool(fNodownload, "noDownload")
	configureBool(fQuiet, "quiet")
	configureBool(fProgress, "progress")

	configureBool(doshPbf, "doshPbf")
	configureBool(dosmPbf, "dosmPbf")
	configureBool(dosmGz, "dosmGz")
	configureBool(dosmBz2, "dosmBz2")
	configureBool(dshpZip, "dshpZip")
	configureBool(dstate, "dstate")
	configureBool(dpoly, "dpoly")
	configureBool(dkml, "dkml")

	viper.Set("service", fService)

	if *dOutputDir == "" {
		if *dOutputDir = os.Getenv("OUTPUT_DIR"); *dOutputDir == "" {
			var err error

			*dOutputDir, err = os.Getwd()
			if err != nil {
				panic(err)
			}
		}
	}

	*dOutputDir += string(os.PathSeparator)

	checkService()

	switch commands {
	case list.FullCommand():
		listCommand()
	case d.FullCommand():
		downloadCommand()
	case generate.FullCommand():
		generator.Generate(*fConfig)
	}
}

func fileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}

	return false
}

func hashFileMD5(filePath string) (string, error) {
	var returnMD5String string

	if fileExist(filePath) {
		hash := md5.New() //nolint:gosec

		file, err := os.Open(filePath)
		if err != nil {
			return returnMD5String, err
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.Panicln(err)
			}
		}()

		if _, err := io.Copy(hash, file); err != nil {
			return returnMD5String, err
		}

		hashInBytes := hash.Sum(nil)[:16]
		returnMD5String = hex.EncodeToString(hashInBytes)

		return returnMD5String, nil
	}

	return returnMD5String, nil
}

func controlHash(hashfile, hash string) (bool, error) {
	if fileExist(hashfile) {
		file, err := ioutil.ReadFile(hashfile)
		if err != nil {
			return false, err
		}

		filehash := strings.Split(string(file), " ")[0]

		if *fVerbose && !*fQuiet {
			log.Println("Hash from file :", filehash)
		}

		return strings.EqualFold(hash, filehash), nil
	}

	return false, nil
}

func downloadChecksum(format string) bool {
	ret := false

	if *dCheck {
		hash := "md5"
		fhash := format + "." + hash

		configPtr, err := config.LoadConfig(*fConfig)
		if err != nil {
			log.Panicln(err)
		}

		if ok, _, _ := config.IsHashable(configPtr, format); ok {
			myElem, err := config.FindElem(configPtr, *delement)
			if err != nil {
				log.Panicln(err)
			}

			myURL, err := config.Elem2URL(configPtr, myElem, fhash)
			if err != nil {
				log.Panicln(err)
			}

			if e := download.FromURL(myURL, *dOutputDir+*delement+"."+fhash); e != nil {
				log.Panicln(e)
			}

			if *fVerbose && !*fQuiet {
				log.Println("Hashing", *dOutputDir+*delement+"."+format)
			}

			hashed, err := hashFileMD5(*dOutputDir + *delement + "." + format)
			if err != nil {
				log.Panicln(err)
			}

			if *fVerbose && !*fQuiet {
				log.Println("MD5 :", hashed)
			}

			ret, err = controlHash(*dOutputDir+*delement+"."+fhash, hashed)
			if err != nil {
				log.Panicln(err)
			}

			if !*fQuiet {
				if ret {
					log.Println("Checksum OK for", *dOutputDir+*delement+"."+format)
				} else {
					log.Println("Checksum MISMATCH for", *dOutputDir+*delement+"."+format)
				}
			}

			return ret
		}

		if !*fQuiet {
			log.Println("No checksum provided for", *dOutputDir+*delement+"."+format)
		}
	}

	return ret
}
