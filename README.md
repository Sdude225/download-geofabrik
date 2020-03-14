# download-geofabrik

[![Build Status](https://travis-ci.org/julien-noblet/download-geofabrik.svg?branch=master)](https://travis-ci.org/julien-noblet/download-geofabrik)
[![Join the chat at https://gitter.im/julien-noblet/download-geofabrik](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/julien-noblet/download-geofabrik?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/julien-noblet/download-geofabrik)](https://goreportcard.com/report/github.com/julien-noblet/download-geofabrik)
[![Coverage Status](https://coveralls.io/repos/github/julien-noblet/download-geofabrik/badge.svg?branch=master)](https://coveralls.io/github/julien-noblet/download-geofabrik?branch=master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik?ref=badge_shield)

## Version 2
Warning! command line have changed from V1
see [Usage](#usage)

## Docker
```shell
docker run -it --rm -v $PWD:/data download-geofabrik:latest download element
```
where ```element``` is one of geofabrik's files.
## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjulien-noblet%2Fdownload-geofabrik?ref=badge_large)

## Usage
```shell
./download-geofabrik download element
```
where ```element``` is one of geofabrik's files.
```shell
./download-geofabrik --help-long

usage: download-geofabrik [<flags>] <command> [<args> ...]

A command-line tool for downloading OSM files.

Flags:
      --help                 Show context-sensitive help (also try --help-long
                             and --help-man).
      --service="geofabrik"  Can switch to another service. You can use
                             "geofabrik", "openstreetmap.fr" or "bbbike". It
                             automatically change config file if -c is unused.
  -c, --config="./geofabrik.yml"  
                             Set Config file.
  -n, --nodownload           Do not download file (test only)
  -v, --verbose              Be verbose
  -q, --quiet                Be quiet
      --progress             Add a progress bar (implie quiet)
      --version              Show application version.

Commands:
  help [<command>...]
    Show help.


  list [<flags>]
    Show elements available

    --markdown  generate list in Markdown format

  download [<flags>] <element>
    Download element

    -B, --osm.bz2  Download osm.bz2 if available
    -G, --osm.gz   Download osm.gz if available
    -S, --shp.zip  Download shp.zip if available
    -P, --osm.pbf  Download osm.pbf (default)
    -H, --osh.pbf  Download osh.pbf
    -s, --state    Download state.txt file
    -p, --poly     Download poly file
    -k, --kml      Download kml file
        --check    Control with checksum (default) Use --no-check to discard
                   control
    -d, --output_directory=OUTPUT_DIRECTORY  
                   Set output directory, you can use also OUTPUT_DIR env
                   variable

  generate
    Generate a new config file



```

## List of elements
|                  SHORTNAME                  | IS IN | LONG NAME | FORMATS |
|---------------------------------------------|-------|-----------|---------|
| afghanistan                                 |       |           | P       |
| africa                                      |       |           | P       |
| albania                                     |       |           | P       |
| alberta                                     |       |           | P       |
| algeria                                     |       |           | P       |
| alps                                        |       |           | P       |
| alsace                                      |       |           | P       |
| andorra                                     |       |           | P       |
| angola                                      |       |           | P       |
| antarctica                                  |       |           | P       |
| aquitaine                                   |       |           | P       |
| argentina                                   |       |           | P       |
| armenia                                     |       |           | P       |
| arnsberg-regbez                             |       |           | P       |
| asia                                        |       |           | P       |
| australia                                   |       |           | P       |
| australia-oceania                           |       |           | P       |
| austria                                     |       |           | P       |
| auvergne                                    |       |           | P       |
| azerbaijan                                  |       |           | P       |
| azores                                      |       |           | P       |
| baden-wuerttemberg                          |       |           | P       |
| bahamas                                     |       |           | P       |
| bangladesh                                  |       |           | P       |
| basse-normandie                             |       |           | P       |
| bayern                                      |       |           | P       |
| bedfordshire                                |       |           | P       |
| belarus                                     |       |           | P       |
| belgium                                     |       |           | P       |
| belize                                      |       |           | P       |
| benin                                       |       |           | P       |
| berkshire                                   |       |           | P       |
| berlin                                      |       |           | P       |
| bhutan                                      |       |           | P       |
| bolivia                                     |       |           | P       |
| bosnia-herzegovina                          |       |           | P       |
| botswana                                    |       |           | P       |
| bourgogne                                   |       |           | P       |
| brandenburg                                 |       |           | P       |
| brazil                                      |       |           | P       |
| bremen                                      |       |           | P       |
| bretagne                                    |       |           | P       |
| bristol                                     |       |           | P       |
| britain-and-ireland                         |       |           | P       |
| british-columbia                            |       |           | P       |
| buckinghamshire                             |       |           | P       |
| bulgaria                                    |       |           | P       |
| burkina-faso                                |       |           | P       |
| burundi                                     |       |           | P       |
| cambodia                                    |       |           | P       |
| cambridgeshire                              |       |           | P       |
| cameroon                                    |       |           | P       |
| canada                                      |       |           | P       |
| canary-islands                              |       |           | P       |
| cape-verde                                  |       |           | P       |
| central-african-republic                    |       |           | P       |
| central-america                             |       |           | P       |
| central-fed-district                        |       |           | P       |
| centre                                      |       |           | P       |
| centro                                      |       |           | P       |
| centro-oeste                                |       |           | P       |
| chad                                        |       |           | P       |
| champagne-ardenne                           |       |           | P       |
| cheshire                                    |       |           | P       |
| chile                                       |       |           | P       |
| china                                       |       |           | P       |
| chubu                                       |       |           | P       |
| chugoku                                     |       |           | P       |
| colombia                                    |       |           | P       |
| comores                                     |       |           | P       |
| congo-brazzaville                           |       |           | P       |
| congo-democratic-republic                   |       |           | P       |
| cornwall                                    |       |           | P       |
| corse                                       |       |           | P       |
| crimean-fed-district                        |       |           | P       |
| croatia                                     |       |           | P       |
| cuba                                        |       |           | P       |
| cumbria                                     |       |           | P       |
| cyprus                                      |       |           | P       |
| czech-republic                              |       |           | P       |
| dach                                        |       |           | P       |
| denmark                                     |       |           | P       |
| derbyshire                                  |       |           | P       |
| detmold-regbez                              |       |           | P       |
| devon                                       |       |           | P       |
| djibouti                                    |       |           | P       |
| dolnoslaskie                                |       |           | P       |
| dorset                                      |       |           | P       |
| drenthe                                     |       |           | P       |
| duesseldorf-regbez                          |       |           | P       |
| durham                                      |       |           | P       |
| east-sussex                                 |       |           | P       |
| east-yorkshire-with-hull                    |       |           | P       |
| ecuador                                     |       |           | P       |
| egypt                                       |       |           | P       |
| enfield                                     |       |           | P       |
| england                                     |       |           | P       |
| equatorial-guinea                           |       |           | P       |
| eritrea                                     |       |           | P       |
| essex                                       |       |           | P       |
| estonia                                     |       |           | P       |
| ethiopia                                    |       |           | P       |
| europe                                      |       |           | P       |
| far-eastern-fed-district                    |       |           | P       |
| faroe-islands                               |       |           | P       |
| fiji                                        |       |           | P       |
| finland                                     |       |           | P       |
| flevoland                                   |       |           | P       |
| france                                      |       |           | P       |
| franche-comte                               |       |           | P       |
| freiburg-regbez                             |       |           | P       |
| friesland                                   |       |           | P       |
| gabon                                       |       |           | P       |
| gcc-states                                  |       |           | P       |
| gelderland                                  |       |           | P       |
| georgia                                     |       |           | P       |
| germany                                     |       |           | P       |
| ghana                                       |       |           | P       |
| gloucestershire                             |       |           | P       |
| great-britain                               |       |           | P       |
| greater-london                              |       |           | P       |
| greater-manchester                          |       |           | P       |
| greece                                      |       |           | P       |
| greenland                                   |       |           | P       |
| groningen                                   |       |           | P       |
| guadeloupe                                  |       |           | P       |
| guatemala                                   |       |           | P       |
| guinea                                      |       |           | P       |
| guinea-bissau                               |       |           | P       |
| guyane                                      |       |           | P       |
| haiti-and-domrep                            |       |           | P       |
| hamburg                                     |       |           | P       |
| hampshire                                   |       |           | P       |
| haute-normandie                             |       |           | P       |
| herefordshire                               |       |           | P       |
| hertfordshire                               |       |           | P       |
| hessen                                      |       |           | P       |
| hokkaido                                    |       |           | P       |
| hungary                                     |       |           | P       |
| iceland                                     |       |           | P       |
| ile-de-france                               |       |           | P       |
| india                                       |       |           | P       |
| indonesia                                   |       |           | P       |
| iran                                        |       |           | P       |
| iraq                                        |       |           | P       |
| ireland-and-northern-ireland                |       |           | P       |
| isle-of-man                                 |       |           | P       |
| isle-of-wight                               |       |           | P       |
| isole                                       |       |           | P       |
| israel-and-palestine                        |       |           | P       |
| italy                                       |       |           | P       |
| ivory-coast                                 |       |           | P       |
| jamaica                                     |       |           | P       |
| japan                                       |       |           | P       |
| jordan                                      |       |           | P       |
| kaliningrad                                 |       |           | P       |
| kansai                                      |       |           | P       |
| kanto                                       |       |           | P       |
| karlsruhe-regbez                            |       |           | P       |
| kazakhstan                                  |       |           | P       |
| kent                                        |       |           | P       |
| kenya                                       |       |           | P       |
| koeln-regbez                                |       |           | P       |
| kosovo                                      |       |           | P       |
| kujawsko-pomorskie                          |       |           | P       |
| kyrgyzstan                                  |       |           | P       |
| kyushu                                      |       |           | P       |
| lancashire                                  |       |           | P       |
| languedoc-roussillon                        |       |           | P       |
| laos                                        |       |           | P       |
| latvia                                      |       |           | P       |
| lebanon                                     |       |           | P       |
| leicestershire                              |       |           | P       |
| lesotho                                     |       |           | P       |
| liberia                                     |       |           | P       |
| libya                                       |       |           | P       |
| liechtenstein                               |       |           | P       |
| limburg                                     |       |           | P       |
| limousin                                    |       |           | P       |
| lincolnshire                                |       |           | P       |
| lithuania                                   |       |           | P       |
| lodzkie                                     |       |           | P       |
| lorraine                                    |       |           | P       |
| lubelskie                                   |       |           | P       |
| lubuskie                                    |       |           | P       |
| luxembourg                                  |       |           | P       |
| macedonia                                   |       |           | P       |
| madagascar                                  |       |           | P       |
| malawi                                      |       |           | P       |
| malaysia-singapore-brunei                   |       |           | P       |
| maldives                                    |       |           | P       |
| mali                                        |       |           | P       |
| malopolskie                                 |       |           | P       |
| malta                                       |       |           | P       |
| manitoba                                    |       |           | P       |
| martinique                                  |       |           | P       |
| mauritania                                  |       |           | P       |
| mauritius                                   |       |           | P       |
| mayotte                                     |       |           | P       |
| mazowieckie                                 |       |           | P       |
| mecklenburg-vorpommern                      |       |           | P       |
| merseyside                                  |       |           | P       |
| mexico                                      |       |           | P       |
| midi-pyrenees                               |       |           | P       |
| mittelfranken                               |       |           | P       |
| moldova                                     |       |           | P       |
| monaco                                      |       |           | P       |
| mongolia                                    |       |           | P       |
| montenegro                                  |       |           | P       |
| morocco                                     |       |           | P       |
| mozambique                                  |       |           | P       |
| muenster-regbez                             |       |           | P       |
| myanmar                                     |       |           | P       |
| namibia                                     |       |           | P       |
| nepal                                       |       |           | P       |
| netherlands                                 |       |           | P       |
| new-brunswick                               |       |           | P       |
| new-caledonia                               |       |           | P       |
| new-zealand                                 |       |           | P       |
| newfoundland-and-labrador                   |       |           | P       |
| nicaragua                                   |       |           | P       |
| niederbayern                                |       |           | P       |
| niedersachsen                               |       |           | P       |
| niger                                       |       |           | P       |
| nigeria                                     |       |           | P       |
| noord-brabant                               |       |           | P       |
| noord-holland                               |       |           | P       |
| norcal                                      |       |           | P       |
| nord-est                                    |       |           | P       |
| nord-ovest                                  |       |           | P       |
| nord-pas-de-calais                          |       |           | P       |
| nordeste                                    |       |           | P       |
| nordrhein-westfalen                         |       |           | P       |
| norfolk                                     |       |           | P       |
| norte                                       |       |           | P       |
| north-america                               |       |           | P       |
| north-caucasus-fed-district                 |       |           | P       |
| north-korea                                 |       |           | P       |
| north-yorkshire                             |       |           | P       |
| northamptonshire                            |       |           | P       |
| northumberland                              |       |           | P       |
| northwest-territories                       |       |           | P       |
| northwestern-fed-district                   |       |           | P       |
| norway                                      |       |           | P       |
| nottinghamshire                             |       |           | P       |
| nova-scotia                                 |       |           | P       |
| nunavut                                     |       |           | P       |
| oberbayern                                  |       |           | P       |
| oberfranken                                 |       |           | P       |
| oberpfalz                                   |       |           | P       |
| ontario                                     |       |           | P       |
| opolskie                                    |       |           | P       |
| overijssel                                  |       |           | P       |
| oxfordshire                                 |       |           | P       |
| pakistan                                    |       |           | P       |
| papua-new-guinea                            |       |           | P       |
| paraguay                                    |       |           | P       |
| pays-de-la-loire                            |       |           | P       |
| peru                                        |       |           | P       |
| philippines                                 |       |           | P       |
| picardie                                    |       |           | P       |
| podkarpackie                                |       |           | P       |
| podlaskie                                   |       |           | P       |
| poitou-charentes                            |       |           | P       |
| poland                                      |       |           | P       |
| pomorskie                                   |       |           | P       |
| portugal                                    |       |           | P       |
| prince-edward-island                        |       |           | P       |
| provence-alpes-cote-d-azur                  |       |           | P       |
| quebec                                      |       |           | P       |
| reunion                                     |       |           | P       |
| rheinland-pfalz                             |       |           | P       |
| rhone-alpes                                 |       |           | P       |
| romania                                     |       |           | P       |
| russia                                      |       |           | P       |
| rutland                                     |       |           | P       |
| rwanda                                      |       |           | P       |
| saarland                                    |       |           | P       |
| sachsen                                     |       |           | P       |
| sachsen-anhalt                              |       |           | P       |
| saint-helena-ascension-and-tristan-da-cunha |       |           | P       |
| sao-tome-and-principe                       |       |           | P       |
| saskatchewan                                |       |           | P       |
| schleswig-holstein                          |       |           | P       |
| schwaben                                    |       |           | P       |
| scotland                                    |       |           | P       |
| senegal-and-gambia                          |       |           | P       |
| serbia                                      |       |           | P       |
| seychelles                                  |       |           | P       |
| shikoku                                     |       |           | P       |
| shropshire                                  |       |           | P       |
| siberian-fed-district                       |       |           | P       |
| sierra-leone                                |       |           | P       |
| slaskie                                     |       |           | P       |
| slovakia                                    |       |           | P       |
| slovenia                                    |       |           | P       |
| socal                                       |       |           | P       |
| somalia                                     |       |           | P       |
| somerset                                    |       |           | P       |
| south-africa                                |       |           | P       |
| south-africa-and-lesotho                    |       |           | P       |
| south-america                               |       |           | P       |
| south-fed-district                          |       |           | P       |
| south-korea                                 |       |           | P       |
| south-sudan                                 |       |           | P       |
| south-yorkshire                             |       |           | P       |
| spain                                       |       |           | P       |
| sri-lanka                                   |       |           | P       |
| staffordshire                               |       |           | P       |
| stuttgart-regbez                            |       |           | P       |
| sud                                         |       |           | P       |
| sudan                                       |       |           | P       |
| sudeste                                     |       |           | P       |
| suffolk                                     |       |           | P       |
| sul                                         |       |           | P       |
| suriname                                    |       |           | P       |
| surrey                                      |       |           | P       |
| swaziland                                   |       |           | P       |
| sweden                                      |       |           | P       |
| swietokrzyskie                              |       |           | P       |
| switzerland                                 |       |           | P       |
| syria                                       |       |           | P       |
| taiwan                                      |       |           | P       |
| tajikistan                                  |       |           | P       |
| tanzania                                    |       |           | P       |
| thailand                                    |       |           | P       |
| thueringen                                  |       |           | P       |
| togo                                        |       |           | P       |
| tohoku                                      |       |           | P       |
| tuebingen-regbez                            |       |           | P       |
| tunisia                                     |       |           | P       |
| turkey                                      |       |           | P       |
| turkmenistan                                |       |           | P       |
| tyne-and-wear                               |       |           | P       |
| uganda                                      |       |           | P       |
| ukraine                                     |       |           | P       |
| unterfranken                                |       |           | P       |
| ural-fed-district                           |       |           | P       |
| uruguay                                     |       |           | P       |
| us                                          |       |           | P       |
| us-midwest                                  |       |           | P       |
| us-northeast                                |       |           | P       |
| us-pacific                                  |       |           | P       |
| us-south                                    |       |           | P       |
| us-west                                     |       |           | P       |
| us/alabama                                  |       |           | P       |
| us/alaska                                   |       |           | P       |
| us/arizona                                  |       |           | P       |
| us/arkansas                                 |       |           | P       |
| us/california                               |       |           | P       |
| us/colorado                                 |       |           | P       |
| us/connecticut                              |       |           | P       |
| us/delaware                                 |       |           | P       |
| us/district-of-columbia                     |       |           | P       |
| us/florida                                  |       |           | P       |
| us/georgia                                  |       |           | P       |
| us/hawaii                                   |       |           | P       |
| us/idaho                                    |       |           | P       |
| us/illinois                                 |       |           | P       |
| us/indiana                                  |       |           | P       |
| us/iowa                                     |       |           | P       |
| us/kansas                                   |       |           | P       |
| us/kentucky                                 |       |           | P       |
| us/louisiana                                |       |           | P       |
| us/maine                                    |       |           | P       |
| us/maryland                                 |       |           | P       |
| us/massachusetts                            |       |           | P       |
| us/michigan                                 |       |           | P       |
| us/minnesota                                |       |           | P       |
| us/mississippi                              |       |           | P       |
| us/missouri                                 |       |           | P       |
| us/montana                                  |       |           | P       |
| us/nebraska                                 |       |           | P       |
| us/nevada                                   |       |           | P       |
| us/new-hampshire                            |       |           | P       |
| us/new-jersey                               |       |           | P       |
| us/new-mexico                               |       |           | P       |
| us/new-york                                 |       |           | P       |
| us/north-carolina                           |       |           | P       |
| us/north-dakota                             |       |           | P       |
| us/ohio                                     |       |           | P       |
| us/oklahoma                                 |       |           | P       |
| us/oregon                                   |       |           | P       |
| us/pennsylvania                             |       |           | P       |
| us/puerto-rico                              |       |           | P       |
| us/rhode-island                             |       |           | P       |
| us/south-carolina                           |       |           | P       |
| us/south-dakota                             |       |           | P       |
| us/tennessee                                |       |           | P       |
| us/texas                                    |       |           | P       |
| us/utah                                     |       |           | P       |
| us/vermont                                  |       |           | P       |
| us/virginia                                 |       |           | P       |
| us/washington                               |       |           | P       |
| us/west-virginia                            |       |           | P       |
| us/wisconsin                                |       |           | P       |
| us/wyoming                                  |       |           | P       |
| utrecht                                     |       |           | P       |
| uzbekistan                                  |       |           | P       |
| venezuela                                   |       |           | P       |
| vietnam                                     |       |           | P       |
| volga-fed-district                          |       |           | P       |
| wales                                       |       |           | P       |
| warminsko-mazurskie                         |       |           | P       |
| warwickshire                                |       |           | P       |
| west-midlands                               |       |           | P       |
| west-sussex                                 |       |           | P       |
| west-yorkshire                              |       |           | P       |
| wielkopolskie                               |       |           | P       |
| wiltshire                                   |       |           | P       |
| worcestershire                              |       |           | P       |
| yemen                                       |       |           | P       |
| yukon                                       |       |           | P       |
| zachodniopomorskie                          |       |           | P       |
| zambia                                      |       |           | P       |
| zeeland                                     |       |           | P       |
| zimbabwe                                    |       |           | P       |
| zuid-holland                                |       |           | P       |
Total elements: 418

## List of elements from openstreetmap.fr
|                SHORTNAME                |              IS IN               |                LONG NAME                | FORMATS |
|-----------------------------------------|----------------------------------|-----------------------------------------|---------|
| abitibi_temiscamingue                   | quebec                           | abitibi_temiscamingue                   | sPp     |
| abruzzo                                 | italy                            | abruzzo                                 | sPp     |
| aceh                                    | indonesia                        | aceh                                    | sPp     |
| acre                                    | north                            | acre                                    | sPp     |
| adygea_republic                         | southern_federal_district        | adygea_republic                         | sPp     |
| afghanistan                             | asia                             | afghanistan                             | sPp     |
| africa                                  |                                  | africa                                  | sPp     |
| ain                                     | rhone_alpes                      | ain                                     | sPp     |
| aisne                                   | picardie                         | aisne                                   | sPp     |
| akershus                                | norway                           | akershus                                | sPp     |
| alagoas                                 | northeast                        | alagoas                                 | sPp     |
| alameda                                 | california                       | alameda                                 | sPp     |
| alberta                                 | canada                           | alberta                                 | sPp     |
| algeria                                 | africa                           | algeria                                 | sPp     |
| allier                                  | auvergne                         | allier                                  | sPp     |
| alpes_de_haute_provence                 | provence_alpes_cote_d_azur       | alpes_de_haute_provence                 | sPp     |
| alpes_maritimes                         | provence_alpes_cote_d_azur       | alpes_maritimes                         | sPp     |
| alpine                                  | california                       | alpine                                  | sPp     |
| alsace                                  | france                           | alsace                                  | sPp     |
| altai_krai                              | siberian_federal_district        | altai_krai                              | sPp     |
| altai_republic                          | siberian_federal_district        | altai_republic                          | sPp     |
| amador                                  | california                       | amador                                  | sPp     |
| amapa                                   | north                            | amapa                                   | sPp     |
| amazonas                                | north                            | amazonas                                | sPp     |
| american_samoa                          | south-america                    | american_samoa                          | sPp     |
| amur_oblast                             | far_eastern_federal_district     | amur_oblast                             | sPp     |
| andalucia                               | spain                            | andalucia                               | sPp     |
| andaman_and_nicobar_islands             | india                            | andaman_and_nicobar_islands             | sPp     |
| andhra_pradesh                          | india                            | andhra_pradesh                          | sPp     |
| angola                                  | africa                           | angola                                  | sPp     |
| anguilla                                | central-america                  | anguilla                                | sPp     |
| anhui                                   | china                            | anhui                                   | sPp     |
| antigua_and_barbuda                     | central-america                  | antigua_and_barbuda                     | sPp     |
| aquitaine                               | france                           | aquitaine                               | sPp     |
| aragon                                  | spain                            | aragon                                  | sPp     |
| ardeche                                 | rhone_alpes                      | ardeche                                 | sPp     |
| ardennes                                | champagne_ardenne                | ardennes                                | sPp     |
| argentina                               | south-america                    | argentina                               | sPp     |
| ariege                                  | midi_pyrenees                    | ariege                                  | sPp     |
| arkhangelsk_oblast                      | northwestern_federal_district    | arkhangelsk_oblast                      | sPp     |
| armenia                                 | asia                             | armenia                                 | sPp     |
| arnsberg                                | nordrhein_westfalen              | arnsberg                                | sPp     |
| aruba                                   | central-america                  | aruba                                   | sPp     |
| arunachal_pradesh                       | india                            | arunachal_pradesh                       | sPp     |
| asia                                    |                                  | asia                                    | sPp     |
| assam                                   | india                            | assam                                   | sPp     |
| astrakhan_oblast                        | southern_federal_district        | astrakhan_oblast                        | sPp     |
| asturias                                | spain                            | asturias                                | sPp     |
| aube                                    | champagne_ardenne                | aube                                    | sPp     |
| aude                                    | languedoc_roussillon             | aude                                    | sPp     |
| aust-agder                              | norway                           | aust-agder                              | sPp     |
| australia                               | oceania                          | australia                               | sPp     |
| australian_capital_territory            | australia                        | australian_capital_territory            | sPp     |
| austria                                 | europe                           | austria                                 | sPp     |
| auvergne                                | france                           | auvergne                                | sPp     |
| aveyron                                 | midi_pyrenees                    | aveyron                                 | sPp     |
| azores                                  | portugal                         | azores                                  | sPp     |
| bahamas                                 | central-america                  | bahamas                                 | sPp     |
| bahia                                   | northeast                        | bahia                                   | sPp     |
| bahrain                                 | asia                             | bahrain                                 | sPp     |
| bali                                    | indonesia                        | bali                                    | sPp     |
| bangka_belitung_islands                 | indonesia                        | bangka_belitung_islands                 | sPp     |
| banskobystricky                         | slovakia                         | banskobystricky                         | sPp     |
| banten                                  | indonesia                        | banten                                  | sPp     |
| barbados                                | central-america                  | barbados                                | sPp     |
| bas_rhin                                | alsace                           | bas_rhin                                | sPp     |
| bas_saint_laurent                       | quebec                           | bas_saint_laurent                       | sPp     |
| bashkortostan_republic                  | volga_federal_district           | bashkortostan_republic                  | sPp     |
| basilicata                              | italy                            | basilicata                              | sPp     |
| basse_normandie                         | france                           | basse_normandie                         | sPp     |
| beijing                                 | china                            | beijing                                 | sPp     |
| belgium                                 | europe                           | belgium                                 | sPp     |
| belgorod_oblast                         | central_federal_district         | belgorod_oblast                         | sPp     |
| bengkulu                                | indonesia                        | bengkulu                                | sPp     |
| benin                                   | africa                           | benin                                   | sPp     |
| bermuda                                 | north-america                    | bermuda                                 | sPp     |
| bhutan                                  | asia                             | bhutan                                  | sPp     |
| bihar                                   | india                            | bihar                                   | sPp     |
| bouches_du_rhone                        | provence_alpes_cote_d_azur       | bouches_du_rhone                        | sPp     |
| bourgogne                               | france                           | bourgogne                               | sPp     |
| bouvet_island                           | africa                           | bouvet_island                           | sPp     |
| bratislavsky                            | slovakia                         | bratislavsky                            | sPp     |
| brazil                                  | south-america                    | brazil                                  | sPp     |
| bretagne                                | france                           | bretagne                                | sPp     |
| british_columbia                        | canada                           | british_columbia                        | sPp     |
| british_indian_ocean_territory          | asia                             | british_indian_ocean_territory          | sPp     |
| british_virgin_islands                  | central-america                  | british_virgin_islands                  | sPp     |
| brunei                                  | asia                             | brunei                                  | sPp     |
| brussels_capital_region                 | belgium                          | brussels_capital_region                 | sPp     |
| bryansk_oblast                          | central_federal_district         | bryansk_oblast                          | sPp     |
| buenos_aires                            | argentina                        | buenos_aires                            | sPp     |
| buenos_aires_city                       | argentina                        | buenos_aires_city                       | sPp     |
| burgenland                              | austria                          | burgenland                              | sPp     |
| burkina_faso                            | africa                           | burkina_faso                            | sPp     |
| burundi                                 | africa                           | burundi                                 | sPp     |
| buryatia_republic                       | siberian_federal_district        | buryatia_republic                       | sPp     |
| buskerud                                | norway                           | buskerud                                | sPp     |
| butte                                   | california                       | butte                                   | sPp     |
| calabria                                | italy                            | calabria                                | sPp     |
| calaveras                               | california                       | calaveras                               | sPp     |
| california                              | us-west                          | california                              | sPp     |
| calvados                                | basse_normandie                  | calvados                                | sPp     |
| cambodia                                | asia                             | cambodia                                | sPp     |
| cameroon                                | africa                           | cameroon                                | sPp     |
| campania                                | italy                            | campania                                | sPp     |
| canada                                  | north-america                    | canada                                  | sPp     |
| canarias                                | spain                            | canarias                                | sPp     |
| cantabria                               | spain                            | cantabria                               | sPp     |
| cantal                                  | auvergne                         | cantal                                  | sPp     |
| cape_verde                              | africa                           | cape_verde                              | sPp     |
| capitale_nationale                      | quebec                           | capitale_nationale                      | sPp     |
| caribbean                               | central-america                  | caribbean                               | sPp     |
| castilla_la_mancha                      | spain                            | castilla_la_mancha                      | sPp     |
| castilla_y_leon                         | spain                            | castilla_y_leon                         | sPp     |
| catalunya                               | spain                            | catalunya                               | sPp     |
| catamarca                               | argentina                        | catamarca                               | sPp     |
| cayman_islands                          | central-america                  | cayman_islands                          | sPp     |
| ceara                                   | northeast                        | ceara                                   | sPp     |
| central-america                         |                                  | central-america                         | sPp     |
| central-west                            | brazil                           | central-west                            | sPp     |
| central_african_republic                | africa                           | central_african_republic                | sPp     |
| central_federal_district                | russia                           | central_federal_district                | sPp     |
| central_java                            | indonesia                        | central_java                            | sPp     |
| central_kalimantan                      | indonesia                        | central_kalimantan                      | sPp     |
| central_ontario                         | ontario                          | central_ontario                         | sPp     |
| central_sulawesi                        | indonesia                        | central_sulawesi                        | sPp     |
| centre                                  | france                           | centre                                  | sPp     |
| centre_du_quebec                        | quebec                           | centre_du_quebec                        | sPp     |
| ceuta                                   | spain                            | ceuta                                   | sPp     |
| chaco                                   | argentina                        | chaco                                   | sPp     |
| chad                                    | africa                           | chad                                    | sPp     |
| champagne_ardenne                       | france                           | champagne_ardenne                       | sPp     |
| chandigarh                              | india                            | chandigarh                              | sPp     |
| charente                                | poitou_charentes                 | charente                                | sPp     |
| charente_maritime                       | poitou_charentes                 | charente_maritime                       | sPp     |
| chaudiere_appalaches                    | quebec                           | chaudiere_appalaches                    | sPp     |
| chechen_republic                        | north_caucasian_federal_district | chechen_republic                        | sPp     |
| chelyabinsk_oblast                      | ural_federal_district            | chelyabinsk_oblast                      | sPp     |
| cher                                    | centre                           | cher                                    | sPp     |
| cherkasy_oblast                         | ukraine                          | cherkasy_oblast                         | sPp     |
| chernihiv_oblast                        | ukraine                          | chernihiv_oblast                        | sPp     |
| chernivtsi_oblast                       | ukraine                          | chernivtsi_oblast                       | sPp     |
| chhattisgarh                            | india                            | chhattisgarh                            | sPp     |
| china                                   | asia                             | china                                   | sPp     |
| chongqing                               | china                            | chongqing                               | sPp     |
| christmas_island                        | australia                        | christmas_island                        | sPp     |
| chubu                                   | japan                            | chubu                                   | sPp     |
| chubut                                  | argentina                        | chubut                                  | sPp     |
| chugoku                                 | japan                            | chugoku                                 | sPp     |
| chukotka_autonomous_okrug               | far_eastern_federal_district     | chukotka_autonomous_okrug               | sPp     |
| chuvash_republic                        | volga_federal_district           | chuvash_republic                        | sPp     |
| cocos_islands                           | australia                        | cocos_islands                           | sPp     |
| colusa                                  | california                       | colusa                                  | sPp     |
| comoros                                 | africa                           | comoros                                 | sPp     |
| comunidad_de_madrid                     | spain                            | comunidad_de_madrid                     | sPp     |
| comunidad_foral_de_navarra              | spain                            | comunidad_foral_de_navarra              | sPp     |
| comunitat_valenciana                    | spain                            | comunitat_valenciana                    | sPp     |
| congo_brazzaville                       | africa                           | congo_brazzaville                       | sPp     |
| congo_kinshasa                          | africa                           | congo_kinshasa                          | sPp     |
| contra_costa                            | california                       | contra_costa                            | sPp     |
| cook_islands                            | south-america                    | cook_islands                            | sPp     |
| coral_sea_islands                       | australia                        | coral_sea_islands                       | sPp     |
| cordoba                                 | argentina                        | cordoba                                 | sPp     |
| correze                                 | limousin                         | correze                                 | sPp     |
| corrientes                              | argentina                        | corrientes                              | sPp     |
| corse                                   | france                           | corse                                   | sPp     |
| corse_du_sud                            | corse                            | corse_du_sud                            | sPp     |
| costa_rica                              | central-america                  | costa_rica                              | sPp     |
| cote_d_or                               | bourgogne                        | cote_d_or                               | sPp     |
| cote_nord                               | quebec                           | cote_nord                               | sPp     |
| cotes_d_armor                           | bretagne                         | cotes_d_armor                           | sPp     |
| creuse                                  | limousin                         | creuse                                  | sPp     |
| crimea                                  | ukraine                          | crimea                                  | sPp     |
| crimea_republic                         | southern_federal_district        | crimea_republic                         | sPp     |
| curacao                                 | central-america                  | curacao                                 | sPp     |
| czech_republic                          | europe                           | czech_republic                          | sPp     |
| dadra_and_nagar_haveli                  | india                            | dadra_and_nagar_haveli                  | sPp     |
| dagestan_republic                       | north_caucasian_federal_district | dagestan_republic                       | sPp     |
| daman_and_diu                           | india                            | daman_and_diu                           | sPp     |
| del_norte                               | california                       | del_norte                               | sPp     |
| detmold                                 | nordrhein_westfalen              | detmold                                 | sPp     |
| deux_sevres                             | poitou_charentes                 | deux_sevres                             | sPp     |
| distrito-federal                        | central-west                     | distrito-federal                        | sPp     |
| djibouti                                | africa                           | djibouti                                | sPp     |
| dnipropetrovsk_oblast                   | ukraine                          | dnipropetrovsk_oblast                   | sPp     |
| dolnoslaskie                            | poland                           | dolnoslaskie                            | sPp     |
| dominica                                | central-america                  | dominica                                | sPp     |
| dominican_republic                      | central-america                  | dominican_republic                      | sPp     |
| donetsk_oblast                          | ukraine                          | donetsk_oblast                          | sPp     |
| dordogne                                | aquitaine                        | dordogne                                | sPp     |
| doubs                                   | franche_comte                    | doubs                                   | sPp     |
| drenthe                                 | netherlands                      | drenthe                                 | sPp     |
| drome                                   | rhone_alpes                      | drome                                   | sPp     |
| dusseldorf                              | nordrhein_westfalen              | dusseldorf                              | sPp     |
| east                                    | england                          | east                                    | sPp     |
| east_java                               | indonesia                        | east_java                               | sPp     |
| east_kalimantan                         | indonesia                        | east_kalimantan                         | sPp     |
| east_midlands                           | england                          | east_midlands                           | sPp     |
| east_nusa_tenggara                      | indonesia                        | east_nusa_tenggara                      | sPp     |
| east_timor                              | asia                             | east_timor                              | sPp     |
| eastern_ontario                         | ontario                          | eastern_ontario                         | sPp     |
| el_dorado                               | california                       | el_dorado                               | sPp     |
| el_salvador                             | central-america                  | el_salvador                             | sPp     |
| emilia_romagna                          | italy                            | emilia_romagna                          | sPp     |
| england                                 | united_kingdom                   | england                                 | sPp     |
| entre_rios                              | argentina                        | entre_rios                              | sPp     |
| equatorial_guinea                       | africa                           | equatorial_guinea                       | sPp     |
| eritrea                                 | africa                           | eritrea                                 | sPp     |
| espirito-santo                          | southeast                        | espirito-santo                          | sPp     |
| essonne                                 | ile_de_france                    | essonne                                 | sPp     |
| estrie                                  | quebec                           | estrie                                  | sPp     |
| eure                                    | haute_normandie                  | eure                                    | sPp     |
| eure_et_loir                            | centre                           | eure_et_loir                            | sPp     |
| europe                                  |                                  | europe                                  | sPp     |
| euskadi                                 | spain                            | euskadi                                 | sPp     |
| extremadura                             | spain                            | extremadura                             | sPp     |
| falkland                                | south-america                    | falkland                                | sPp     |
| far_eastern_federal_district            | russia                           | far_eastern_federal_district            | sPp     |
| fiji                                    | merge                            | fiji                                    | sP      |
| fiji_east                               | south-america                    | fiji_east                               | sPp     |
| fiji_west                               | oceania                          | fiji_west                               | sPp     |
| finistere                               | bretagne                         | finistere                               | sPp     |
| finnmark                                | norway                           | finnmark                                | sPp     |
| flanders                                | belgium                          | flanders                                | sPp     |
| flevoland                               | netherlands                      | flevoland                               | sPp     |
| formosa                                 | argentina                        | formosa                                 | sPp     |
| france                                  | europe                           | france                                  | sPp     |
| france_metro_dom_com_nc                 | merge                            | france_metro_dom_com_nc                 | sP      |
| france_taaf                             | oceania                          | france_taaf                             | sPp     |
| franche_comte                           | france                           | franche_comte                           | sPp     |
| fresno                                  | california                       | fresno                                  | sPp     |
| friesland                               | netherlands                      | friesland                               | sPp     |
| friuli_venezia_giulia                   | italy                            | friuli_venezia_giulia                   | sPp     |
| fujian                                  | china                            | fujian                                  | sPp     |
| gabon                                   | africa                           | gabon                                   | sPp     |
| galicia                                 | spain                            | galicia                                 | sPp     |
| gambia                                  | africa                           | gambia                                  | sPp     |
| gansu                                   | china                            | gansu                                   | sPp     |
| gard                                    | languedoc_roussillon             | gard                                    | sPp     |
| gaspesie_iles_de_la_madeleine           | quebec                           | gaspesie_iles_de_la_madeleine           | sPp     |
| gelderland                              | netherlands                      | gelderland                              | sPp     |
| georgia                                 | asia                             | georgia                                 | sPp     |
| germany                                 | europe                           | germany                                 | sPp     |
| gers                                    | midi_pyrenees                    | gers                                    | sPp     |
| ghana                                   | africa                           | ghana                                   | sPp     |
| gibraltar                               | europe                           | gibraltar                               | sPp     |
| gironde                                 | aquitaine                        | gironde                                 | sPp     |
| glenn                                   | california                       | glenn                                   | sPp     |
| goa                                     | india                            | goa                                     | sPp     |
| goias                                   | central-west                     | goias                                   | sPp     |
| golden_horseshoe                        | ontario                          | golden_horseshoe                        | sPp     |
| gorontalo                               | indonesia                        | gorontalo                               | sPp     |
| greater_london                          | england                          | greater_london                          | sPp     |
| grenada                                 | central-america                  | grenada                                 | sPp     |
| groningen                               | netherlands                      | groningen                               | sPp     |
| guadeloupe                              | central-america                  | guadeloupe                              | sPp     |
| guam                                    | oceania                          | guam                                    | sPp     |
| guangdong                               | china                            | guangdong                               | sPp     |
| guangxi                                 | china                            | guangxi                                 | sPp     |
| guernesey                               | europe                           | guernesey                               | sPp     |
| guinea                                  | africa                           | guinea                                  | sPp     |
| guizhou                                 | china                            | guizhou                                 | sPp     |
| gujarat                                 | india                            | gujarat                                 | sPp     |
| guyana                                  | south-america                    | guyana                                  | sPp     |
| guyane                                  | south-america                    | guyane                                  | sPp     |
| hainan                                  | china                            | hainan                                  | sPp     |
| haryana                                 | india                            | haryana                                 | sPp     |
| haut_rhin                               | alsace                           | haut_rhin                               | sPp     |
| haute_corse                             | corse                            | haute_corse                             | sPp     |
| haute_garonne                           | midi_pyrenees                    | haute_garonne                           | sPp     |
| haute_loire                             | auvergne                         | haute_loire                             | sPp     |
| haute_marne                             | champagne_ardenne                | haute_marne                             | sPp     |
| haute_normandie                         | france                           | haute_normandie                         | sPp     |
| haute_saone                             | franche_comte                    | haute_saone                             | sPp     |
| haute_savoie                            | rhone_alpes                      | haute_savoie                            | sPp     |
| haute_vienne                            | limousin                         | haute_vienne                            | sPp     |
| hautes_alpes                            | provence_alpes_cote_d_azur       | hautes_alpes                            | sPp     |
| hautes_pyrenees                         | midi_pyrenees                    | hautes_pyrenees                         | sPp     |
| hauts_de_seine                          | ile_de_france                    | hauts_de_seine                          | sPp     |
| hebei                                   | china                            | hebei                                   | sPp     |
| hedmark                                 | norway                           | hedmark                                 | sPp     |
| heilongjiang                            | china                            | heilongjiang                            | sPp     |
| henan                                   | china                            | henan                                   | sPp     |
| herault                                 | languedoc_roussillon             | herault                                 | sPp     |
| himachal_pradesh                        | india                            | himachal_pradesh                        | sPp     |
| hokkaido                                | japan                            | hokkaido                                | sPp     |
| honduras                                | central-america                  | honduras                                | sPp     |
| hong_kong                               | china                            | hong_kong                               | sPp     |
| hordaland                               | norway                           | hordaland                               | sPp     |
| hubei                                   | china                            | hubei                                   | sPp     |
| humboldt                                | california                       | humboldt                                | sPp     |
| hunan                                   | china                            | hunan                                   | sPp     |
| ile_de_france                           | france                           | ile_de_france                           | sPp     |
| ille_et_vilaine                         | bretagne                         | ille_et_vilaine                         | sPp     |
| illes_balears                           | spain                            | illes_balears                           | sPp     |
| imperial                                | california                       | imperial                                | sPp     |
| india                                   | asia                             | india                                   | sPp     |
| indonesia                               | asia                             | indonesia                               | sPp     |
| indre                                   | centre                           | indre                                   | sPp     |
| indre_et_loire                          | centre                           | indre_et_loire                          | sPp     |
| ingushetia_republic                     | north_caucasian_federal_district | ingushetia_republic                     | sPp     |
| inner_mongolia                          | china                            | inner_mongolia                          | sPp     |
| inyo                                    | california                       | inyo                                    | sPp     |
| ionian_sea                              | seas                             | ionian_sea                              | sPp     |
| ireland                                 | europe                           | ireland                                 | sPp     |
| irkutsk_oblast                          | siberian_federal_district        | irkutsk_oblast                          | sPp     |
| isere                                   | rhone_alpes                      | isere                                   | sPp     |
| israel                                  | asia                             | israel                                  | sPp     |
| israel_and_palestine                    | asia                             | israel_and_palestine                    | sPp     |
| italy                                   | europe                           | italy                                   | sPp     |
| ivano-frankivsk_oblast                  | ukraine                          | ivano-frankivsk_oblast                  | sPp     |
| ivanovo_oblast                          | central_federal_district         | ivanovo_oblast                          | sPp     |
| ivory_coast                             | africa                           | ivory_coast                             | sPp     |
| jakarta                                 | indonesia                        | jakarta                                 | sPp     |
| jamaica                                 | central-america                  | jamaica                                 | sPp     |
| jambi                                   | indonesia                        | jambi                                   | sPp     |
| jammu_and_kashmir                       | india                            | jammu_and_kashmir                       | sPp     |
| jan_mayen                               | norway                           | jan_mayen                               | sPp     |
| japan                                   | asia                             | japan                                   | sPp     |
| jersey                                  | europe                           | jersey                                  | sPp     |
| jewish_autonomous_oblast                | far_eastern_federal_district     | jewish_autonomous_oblast                | sPp     |
| jharkhand                               | india                            | jharkhand                               | sPp     |
| jiangsu                                 | china                            | jiangsu                                 | sPp     |
| jiangxi                                 | china                            | jiangxi                                 | sPp     |
| jihocesky                               | czech_republic                   | jihocesky                               | sPp     |
| jihomoravsky                            | czech_republic                   | jihomoravsky                            | sPp     |
| jilin                                   | china                            | jilin                                   | sPp     |
| jujuy                                   | argentina                        | jujuy                                   | sPp     |
| jura                                    | franche_comte                    | jura                                    | sPp     |
| kabardino_balkar_republic               | north_caucasian_federal_district | kabardino_balkar_republic               | sPp     |
| kaliningrad_oblast                      | northwestern_federal_district    | kaliningrad_oblast                      | sPp     |
| kalmykia_republic                       | southern_federal_district        | kalmykia_republic                       | sPp     |
| kaluga_oblast                           | central_federal_district         | kaluga_oblast                           | sPp     |
| kamchatka_krai                          | far_eastern_federal_district     | kamchatka_krai                          | sPp     |
| kansai                                  | japan                            | kansai                                  | sPp     |
| kanto                                   | japan                            | kanto                                   | sPp     |
| karachay_cherkess_republic              | north_caucasian_federal_district | karachay_cherkess_republic              | sPp     |
| karelia_republic                        | northwestern_federal_district    | karelia_republic                        | sPp     |
| karlovarsky                             | czech_republic                   | karlovarsky                             | sPp     |
| karnataka                               | india                            | karnataka                               | sPp     |
| karnten                                 | austria                          | karnten                                 | sPp     |
| kemerovo_oblast                         | siberian_federal_district        | kemerovo_oblast                         | sPp     |
| kenya                                   | africa                           | kenya                                   | sPp     |
| kerala                                  | india                            | kerala                                  | sPp     |
| kern                                    | california                       | kern                                    | sPp     |
| khabarovsk_krai                         | far_eastern_federal_district     | khabarovsk_krai                         | sPp     |
| khakassia_republic                      | siberian_federal_district        | khakassia_republic                      | sPp     |
| khanty_mansi_autonomous_okrug           | ural_federal_district            | khanty_mansi_autonomous_okrug           | sPp     |
| kharkiv_oblast                          | ukraine                          | kharkiv_oblast                          | sPp     |
| kherson_oblast                          | ukraine                          | kherson_oblast                          | sPp     |
| khmelnytskyi_oblast                     | ukraine                          | khmelnytskyi_oblast                     | sPp     |
| kiev                                    | ukraine                          | kiev                                    | sPp     |
| kiev_oblast                             | ukraine                          | kiev_oblast                             | sPp     |
| kings                                   | california                       | kings                                   | sPp     |
| kiribati                                | merge                            | kiribati                                | sP      |
| kiribati_east                           | south-america                    | kiribati_east                           | sPp     |
| kiribati_west                           | oceania                          | kiribati_west                           | sPp     |
| kirov_oblast                            | volga_federal_district           | kirov_oblast                            | sPp     |
| kirovohrad_oblast                       | ukraine                          | kirovohrad_oblast                       | sPp     |
| koln                                    | nordrhein_westfalen              | koln                                    | sPp     |
| komi_republic                           | northwestern_federal_district    | komi_republic                           | sPp     |
| kosicky                                 | slovakia                         | kosicky                                 | sPp     |
| kostroma_oblast                         | central_federal_district         | kostroma_oblast                         | sPp     |
| kralovehradecky                         | czech_republic                   | kralovehradecky                         | sPp     |
| krasnodar_krai                          | southern_federal_district        | krasnodar_krai                          | sPp     |
| krasnoyarsk_krai                        | siberian_federal_district        | krasnoyarsk_krai                        | sPp     |
| kujawsko_pomorskie                      | poland                           | kujawsko_pomorskie                      | sPp     |
| kurgan_oblast                           | ural_federal_district            | kurgan_oblast                           | sPp     |
| kursk_oblast                            | central_federal_district         | kursk_oblast                            | sPp     |
| kuwait                                  | asia                             | kuwait                                  | sPp     |
| kyushu                                  | japan                            | kyushu                                  | sPp     |
| la_pampa                                | argentina                        | la_pampa                                | sPp     |
| la_rioja                                | argentina                        | la_rioja                                | sPp     |
| lake                                    | california                       | lake                                    | sPp     |
| lakshadweep                             | india                            | lakshadweep                             | sPp     |
| lampung                                 | indonesia                        | lampung                                 | sPp     |
| lanaudiere                              | quebec                           | lanaudiere                              | sPp     |
| landes                                  | aquitaine                        | landes                                  | sPp     |
| languedoc_roussillon                    | france                           | languedoc_roussillon                    | sPp     |
| laos                                    | asia                             | laos                                    | sPp     |
| lassen                                  | california                       | lassen                                  | sPp     |
| laurentides                             | quebec                           | laurentides                             | sPp     |
| laval                                   | quebec                           | laval                                   | sPp     |
| lazio                                   | italy                            | lazio                                   | sPp     |
| leningrad_oblast                        | northwestern_federal_district    | leningrad_oblast                        | sPp     |
| lesotho                                 | africa                           | lesotho                                 | sPp     |
| liaoning                                | china                            | liaoning                                | sPp     |
| liberecky                               | czech_republic                   | liberecky                               | sPp     |
| liguria                                 | italy                            | liguria                                 | sPp     |
| limburg                                 | netherlands                      | limburg                                 | sPp     |
| limousin                                | france                           | limousin                                | sPp     |
| lipetsk_oblast                          | central_federal_district         | lipetsk_oblast                          | sPp     |
| lodzkie                                 | poland                           | lodzkie                                 | sPp     |
| loir_et_cher                            | centre                           | loir_et_cher                            | sPp     |
| loire                                   | rhone_alpes                      | loire                                   | sPp     |
| loire_atlantique                        | pays_de_la_loire                 | loire_atlantique                        | sPp     |
| loiret                                  | centre                           | loiret                                  | sPp     |
| lombardia                               | italy                            | lombardia                               | sPp     |
| lorraine                                | france                           | lorraine                                | sPp     |
| los_angeles                             | california                       | los_angeles                             | sPp     |
| lot                                     | midi_pyrenees                    | lot                                     | sPp     |
| lot_et_garonne                          | aquitaine                        | lot_et_garonne                          | sPp     |
| lozere                                  | languedoc_roussillon             | lozere                                  | sPp     |
| lubelskie                               | poland                           | lubelskie                               | sPp     |
| lubuskie                                | poland                           | lubuskie                                | sPp     |
| luhansk_oblast                          | ukraine                          | luhansk_oblast                          | sPp     |
| luxembourg                              | europe                           | luxembourg                              | sPp     |
| lviv_oblast                             | ukraine                          | lviv_oblast                             | sPp     |
| macau                                   | china                            | macau                                   | sPp     |
| madeira                                 | portugal                         | madeira                                 | sPp     |
| madera                                  | california                       | madera                                  | sPp     |
| madhya_pradesh                          | india                            | madhya_pradesh                          | sPp     |
| magadan_oblast                          | far_eastern_federal_district     | magadan_oblast                          | sPp     |
| maharashtra                             | india                            | maharashtra                             | sPp     |
| maine_et_loire                          | pays_de_la_loire                 | maine_et_loire                          | sPp     |
| malawi                                  | africa                           | malawi                                  | sPp     |
| malaysia                                | asia                             | malaysia                                | sPp     |
| maldives                                | asia                             | maldives                                | sPp     |
| mali                                    | africa                           | mali                                    | sPp     |
| malopolskie                             | poland                           | malopolskie                             | sPp     |
| maluku                                  | indonesia                        | maluku                                  | sPp     |
| manche                                  | basse_normandie                  | manche                                  | sPp     |
| manipur                                 | india                            | manipur                                 | sPp     |
| manitoba                                | canada                           | manitoba                                | sPp     |
| maranhao                                | northeast                        | maranhao                                | sPp     |
| marche                                  | italy                            | marche                                  | sPp     |
| mari_el_republic                        | volga_federal_district           | mari_el_republic                        | sPp     |
| marin                                   | california                       | marin                                   | sPp     |
| mariposa                                | california                       | mariposa                                | sPp     |
| marne                                   | champagne_ardenne                | marne                                   | sPp     |
| marshall-islands                        | south-america                    | marshall-islands                        | sPp     |
| marshall_islands                        | oceania                          | marshall_islands                        | sPp     |
| martinique                              | central-america                  | martinique                              | sPp     |
| mato-grosso                             | central-west                     | mato-grosso                             | sPp     |
| mato-grosso-do-sul                      | central-west                     | mato-grosso-do-sul                      | sPp     |
| mauricie                                | quebec                           | mauricie                                | sPp     |
| mauritania                              | africa                           | mauritania                              | sPp     |
| mauritius                               | africa                           | mauritius                               | sPp     |
| mayenne                                 | pays_de_la_loire                 | mayenne                                 | sPp     |
| mayotte                                 | africa                           | mayotte                                 | sPp     |
| mazowieckie                             | poland                           | mazowieckie                             | sPp     |
| meghalaya                               | india                            | meghalaya                               | sPp     |
| melilla                                 | spain                            | melilla                                 | sPp     |
| mendocino                               | california                       | mendocino                               | sPp     |
| mendoza                                 | argentina                        | mendoza                                 | sPp     |
| merced                                  | california                       | merced                                  | sPp     |
| merge                                   |                                  | merge                                   |         |
| meurthe_et_moselle                      | lorraine                         | meurthe_et_moselle                      | sPp     |
| meuse                                   | lorraine                         | meuse                                   | sPp     |
| micronesia                              | oceania                          | micronesia                              | sPp     |
| midi_pyrenees                           | france                           | midi_pyrenees                           | sPp     |
| minas-gerais                            | southeast                        | minas-gerais                            | sPp     |
| misiones                                | argentina                        | misiones                                | sPp     |
| mizoram                                 | india                            | mizoram                                 | sPp     |
| modoc                                   | california                       | modoc                                   | sPp     |
| moere_og_romsdal                        | norway                           | moere_og_romsdal                        | sPp     |
| molise                                  | italy                            | molise                                  | sPp     |
| monaco                                  | europe                           | monaco                                  | sPp     |
| mono                                    | california                       | mono                                    | sPp     |
| monteregie                              | quebec                           | monteregie                              | sPp     |
| monterey                                | california                       | monterey                                | sPp     |
| montreal                                | quebec                           | montreal                                | sPp     |
| montserrat                              | central-america                  | montserrat                              | sPp     |
| moravskoslezsky                         | czech_republic                   | moravskoslezsky                         | sPp     |
| morbihan                                | bretagne                         | morbihan                                | sPp     |
| mordovia_republic                       | volga_federal_district           | mordovia_republic                       | sPp     |
| moscow                                  | central_federal_district         | moscow                                  | sPp     |
| moscow_oblast                           | central_federal_district         | moscow_oblast                           | sPp     |
| moselle                                 | lorraine                         | moselle                                 | sPp     |
| mozambique                              | africa                           | mozambique                              | sPp     |
| munster                                 | nordrhein_westfalen              | munster                                 | sPp     |
| murmansk_oblast                         | northwestern_federal_district    | murmansk_oblast                         | sPp     |
| myanmar                                 | asia                             | myanmar                                 | sPp     |
| mykolaiv_oblast                         | ukraine                          | mykolaiv_oblast                         | sPp     |
| nagaland                                | india                            | nagaland                                | sPp     |
| namibia                                 | africa                           | namibia                                 | sPp     |
| napa                                    | california                       | napa                                    | sPp     |
| national_capital_territory_of_delhi     | india                            | national_capital_territory_of_delhi     | sPp     |
| nauri                                   | south-america                    | nauri                                   | sPp     |
| nauru                                   | oceania                          | nauru                                   | sPp     |
| nenets_autonomous_okrug                 | northwestern_federal_district    | nenets_autonomous_okrug                 | sPp     |
| netherlands                             | europe                           | netherlands                             | sPp     |
| neuquen                                 | argentina                        | neuquen                                 | sPp     |
| nevada                                  | california                       | nevada                                  | sPp     |
| new_brunswick                           | canada                           | new_brunswick                           | sPp     |
| new_caledonia                           | oceania                          | new_caledonia                           | sPp     |
| new_south_wales                         | australia                        | new_south_wales                         | sPp     |
| newfoundland_and_labrador               | canada                           | newfoundland_and_labrador               | sPp     |
| nicaragua                               | central-america                  | nicaragua                               | sPp     |
| niederosterreich                        | austria                          | niederosterreich                        | sPp     |
| nievre                                  | bourgogne                        | nievre                                  | sPp     |
| niger                                   | africa                           | niger                                   | sPp     |
| ningxia                                 | china                            | ningxia                                 | sPp     |
| nitriansky                              | slovakia                         | nitriansky                              | sPp     |
| niue                                    | south-america                    | niue                                    | sPp     |
| nizhny_novgorod_oblast                  | volga_federal_district           | nizhny_novgorod_oblast                  | sPp     |
| noord_brabant                           | netherlands                      | noord_brabant                           | sPp     |
| noord_holland                           | netherlands                      | noord_holland                           | sPp     |
| nord                                    | nord_pas_de_calais               | nord                                    | sPp     |
| nord_du_quebec                          | quebec                           | nord_du_quebec                          | sPp     |
| nord_pas_de_calais                      | france                           | nord_pas_de_calais                      | sPp     |
| nordland                                | norway                           | nordland                                | sPp     |
| nordrhein_westfalen                     | germany                          | nordrhein_westfalen                     | sPp     |
| norfolk_island                          | australia                        | norfolk_island                          | sPp     |
| north                                   | brazil                           | north                                   | sPp     |
| north-america                           |                                  | north-america                           | sPp     |
| north_caucasian_federal_district        | russia                           | north_caucasian_federal_district        | sPp     |
| north_east                              | england                          | north_east                              | sPp     |
| north_kalimantan                        | indonesia                        | north_kalimantan                        | sPp     |
| north_maluku                            | indonesia                        | north_maluku                            | sPp     |
| north_ossetia_alania_republic           | north_caucasian_federal_district | north_ossetia_alania_republic           | sPp     |
| north_sea                               | seas                             | north_sea                               | sPp     |
| north_sulawesi                          | indonesia                        | north_sulawesi                          | sPp     |
| north_sumatra                           | indonesia                        | north_sumatra                           | sPp     |
| north_west                              | england                          | north_west                              | sPp     |
| northeast                               | brazil                           | northeast                               | sPp     |
| northeastern_ontario                    | ontario                          | northeastern_ontario                    | sPp     |
| northern_ireland                        | united_kingdom                   | northern_ireland                        | sPp     |
| northern_mariana_islands                | oceania                          | northern_mariana_islands                | sPp     |
| northern_territory                      | australia                        | northern_territory                      | sPp     |
| northwest_territories                   | canada                           | northwest_territories                   | sPp     |
| northwestern_federal_district           | russia                           | northwestern_federal_district           | sPp     |
| northwestern_ontario                    | ontario                          | northwestern_ontario                    | sPp     |
| norway                                  | europe                           | norway                                  | sPp     |
| nova_scotia                             | canada                           | nova_scotia                             | sPp     |
| novgorod_oblast                         | northwestern_federal_district    | novgorod_oblast                         | sPp     |
| novosibirsk_oblast                      | siberian_federal_district        | novosibirsk_oblast                      | sPp     |
| nunavut                                 | canada                           | nunavut                                 | sPp     |
| oberosterreich                          | austria                          | oberosterreich                          | sPp     |
| oceania                                 |                                  | oceania                                 | sPp     |
| odessa_oblast                           | ukraine                          | odessa_oblast                           | sPp     |
| odisha                                  | india                            | odisha                                  | sPp     |
| oestfold                                | norway                           | oestfold                                | sPp     |
| oise                                    | picardie                         | oise                                    | sPp     |
| olomoucky                               | czech_republic                   | olomoucky                               | sPp     |
| oman                                    | asia                             | oman                                    | sPp     |
| omsk_oblast                             | siberian_federal_district        | omsk_oblast                             | sPp     |
| ontario                                 | canada                           | ontario                                 | sPp     |
| opolskie                                | poland                           | opolskie                                | sPp     |
| oppland                                 | norway                           | oppland                                 | sPp     |
| orange                                  | california                       | orange                                  | sPp     |
| orenburg_oblast                         | volga_federal_district           | orenburg_oblast                         | sPp     |
| orne                                    | basse_normandie                  | orne                                    | sPp     |
| oryol_oblast                            | central_federal_district         | oryol_oblast                            | sPp     |
| oslo                                    | norway                           | oslo                                    | sPp     |
| outaouais                               | quebec                           | outaouais                               | sPp     |
| overijssel                              | netherlands                      | overijssel                              | sPp     |
| palau                                   | oceania                          | palau                                   | sPp     |
| palestine                               | asia                             | palestine                               | sPp     |
| panama                                  | central-america                  | panama                                  | sPp     |
| papua                                   | indonesia                        | papua                                   | sPp     |
| papua_new_guinea                        | oceania                          | papua_new_guinea                        | sPp     |
| para                                    | north                            | para                                    | sPp     |
| paraguay                                | south-america                    | paraguay                                | sPp     |
| paraiba                                 | northeast                        | paraiba                                 | sPp     |
| parana                                  | south                            | parana                                  | sPp     |
| pardubicky                              | czech_republic                   | pardubicky                              | sPp     |
| paris                                   | ile_de_france                    | paris                                   | sPp     |
| pas_de_calais                           | nord_pas_de_calais               | pas_de_calais                           | sPp     |
| pays_de_la_loire                        | france                           | pays_de_la_loire                        | sPp     |
| penza_oblast                            | volga_federal_district           | penza_oblast                            | sPp     |
| perm_krai                               | volga_federal_district           | perm_krai                               | sPp     |
| pernambuco                              | northeast                        | pernambuco                              | sPp     |
| piaui                                   | northeast                        | piaui                                   | sPp     |
| picardie                                | france                           | picardie                                | sPp     |
| piemonte                                | italy                            | piemonte                                | sPp     |
| pitcairn                                | south-america                    | pitcairn                                | sPp     |
| placer                                  | california                       | placer                                  | sPp     |
| plumas                                  | california                       | plumas                                  | sPp     |
| plzensky                                | czech_republic                   | plzensky                                | sPp     |
| podkarpackie                            | poland                           | podkarpackie                            | sPp     |
| podlaskie                               | poland                           | podlaskie                               | sPp     |
| poitou_charentes                        | france                           | poitou_charentes                        | sPp     |
| poland                                  | europe                           | poland                                  | sPp     |
| poltava_oblast                          | ukraine                          | poltava_oblast                          | sPp     |
| polynesie                               | south-america                    | polynesie                               | sPp     |
| pomorskie                               | poland                           | pomorskie                               | sPp     |
| portugal                                | europe                           | portugal                                | sPp     |
| praha                                   | czech_republic                   | praha                                   | sPp     |
| presovsky                               | slovakia                         | presovsky                               | sPp     |
| primorsky_krai                          | far_eastern_federal_district     | primorsky_krai                          | sPp     |
| prince_edward_island                    | canada                           | prince_edward_island                    | sPp     |
| provence_alpes_cote_d_azur              | france                           | provence_alpes_cote_d_azur              | sPp     |
| pskov_oblast                            | northwestern_federal_district    | pskov_oblast                            | sPp     |
| puducherry                              | india                            | puducherry                              | sPp     |
| puerto_rico                             | central-america                  | puerto_rico                             | sPp     |
| puglia                                  | italy                            | puglia                                  | sPp     |
| punjab                                  | india                            | punjab                                  | sPp     |
| puy_de_dome                             | auvergne                         | puy_de_dome                             | sPp     |
| pyrenees_atlantiques                    | aquitaine                        | pyrenees_atlantiques                    | sPp     |
| pyrenees_orientales                     | languedoc_roussillon             | pyrenees_orientales                     | sPp     |
| qatar                                   | asia                             | qatar                                   | sPp     |
| qinghai                                 | china                            | qinghai                                 | sPp     |
| quebec                                  | canada                           | quebec                                  | sPp     |
| queensland                              | australia                        | queensland                              | sPp     |
| rajasthan                               | india                            | rajasthan                               | sPp     |
| region_de_murcia                        | spain                            | region_de_murcia                        | sPp     |
| reunion                                 | africa                           | reunion                                 | sPp     |
| rhone                                   | rhone_alpes                      | rhone                                   | sPp     |
| rhone_alpes                             | france                           | rhone_alpes                             | sPp     |
| riau                                    | indonesia                        | riau                                    | sPp     |
| riau_islands                            | indonesia                        | riau_islands                            | sPp     |
| rio-de-janeiro                          | southeast                        | rio-de-janeiro                          | sPp     |
| rio-grande-do-norte                     | northeast                        | rio-grande-do-norte                     | sPp     |
| rio-grande-do-sul                       | south                            | rio-grande-do-sul                       | sPp     |
| rio_negro                               | argentina                        | rio_negro                               | sPp     |
| riverside                               | california                       | riverside                               | sPp     |
| rivne_oblast                            | ukraine                          | rivne_oblast                            | sPp     |
| robots                                  |                                  | robots                                  |         |
| rogaland                                | norway                           | rogaland                                | sPp     |
| rondonia                                | north                            | rondonia                                | sPp     |
| roraima                                 | north                            | roraima                                 | sPp     |
| rostov_oblast                           | southern_federal_district        | rostov_oblast                           | sPp     |
| russia                                  |                                  | russia                                  | sPp     |
| rwanda                                  | africa                           | rwanda                                  | sPp     |
| ryazan_oblast                           | central_federal_district         | ryazan_oblast                           | sPp     |
| sacramento                              | california                       | sacramento                              | sPp     |
| saguenay_lac_saint_jean                 | quebec                           | saguenay_lac_saint_jean                 | sPp     |
| saint_barthelemy                        | central-america                  | saint_barthelemy                        | sPp     |
| saint_helena_ascension_tristan_da_cunha | africa                           | saint_helena_ascension_tristan_da_cunha | sPp     |
| saint_kitts_and_nevis                   | central-america                  | saint_kitts_and_nevis                   | sPp     |
| saint_lucia                             | central-america                  | saint_lucia                             | sPp     |
| saint_martin                            | central-america                  | saint_martin                            | sPp     |
| saint_petersburg                        | northwestern_federal_district    | saint_petersburg                        | sPp     |
| saint_pierre_et_miquelon                | north-america                    | saint_pierre_et_miquelon                | sPp     |
| saint_vincent_and_the_grenadines        | central-america                  | saint_vincent_and_the_grenadines        | sPp     |
| sakha_republic                          | far_eastern_federal_district     | sakha_republic                          | sPp     |
| sakhalin_oblast                         | far_eastern_federal_district     | sakhalin_oblast                         | sPp     |
| salta                                   | argentina                        | salta                                   | sPp     |
| salzburg                                | austria                          | salzburg                                | sPp     |
| samara_oblast                           | volga_federal_district           | samara_oblast                           | sPp     |
| samoa                                   | south-america                    | samoa                                   | sPp     |
| san_benito                              | california                       | san_benito                              | sPp     |
| san_bernardino                          | california                       | san_bernardino                          | sPp     |
| san_diego                               | california                       | san_diego                               | sPp     |
| san_francisco                           | california                       | san_francisco                           | sPp     |
| san_joaquin                             | california                       | san_joaquin                             | sPp     |
| san_juan                                | argentina                        | san_juan                                | sPp     |
| san_luis                                | argentina                        | san_luis                                | sPp     |
| san_luis_obispo                         | california                       | san_luis_obispo                         | sPp     |
| san_mateo                               | california                       | san_mateo                               | sPp     |
| santa-catarina                          | south                            | santa-catarina                          | sPp     |
| santa_barbara                           | california                       | santa_barbara                           | sPp     |
| santa_clara                             | california                       | santa_clara                             | sPp     |
| santa_cruz                              | california                       | santa_cruz                              | sPp     |
| santa_fe                                | argentina                        | santa_fe                                | sPp     |
| santiago_del_estero                     | argentina                        | santiago_del_estero                     | sPp     |
| sao-paulo                               | southeast                        | sao-paulo                               | sPp     |
| sao_tome_and_principe                   | africa                           | sao_tome_and_principe                   | sPp     |
| saone_et_loire                          | bourgogne                        | saone_et_loire                          | sPp     |
| saratov_oblast                          | volga_federal_district           | saratov_oblast                          | sPp     |
| sardegna                                | italy                            | sardegna                                | sPp     |
| sarthe                                  | pays_de_la_loire                 | sarthe                                  | sPp     |
| saskatchewan                            | canada                           | saskatchewan                            | sPp     |
| saudi_arabia                            | asia                             | saudi_arabia                            | sPp     |
| savoie                                  | rhone_alpes                      | savoie                                  | sPp     |
| seas                                    | europe                           | seas                                    |         |
| seine_et_marne                          | ile_de_france                    | seine_et_marne                          | sPp     |
| seine_maritime                          | haute_normandie                  | seine_maritime                          | sPp     |
| seine_saint_denis                       | ile_de_france                    | seine_saint_denis                       | sPp     |
| senegal                                 | africa                           | senegal                                 | sPp     |
| sergipe                                 | northeast                        | sergipe                                 | sPp     |
| sevastopol                              | southern_federal_district        | sevastopol                              | sPp     |
| seychelles                              | africa                           | seychelles                              | sPp     |
| shaanxi                                 | china                            | shaanxi                                 | sPp     |
| shandong                                | china                            | shandong                                | sPp     |
| shanghai                                | china                            | shanghai                                | sPp     |
| shanxi                                  | china                            | shanxi                                  | sPp     |
| shasta                                  | california                       | shasta                                  | sPp     |
| shikoku                                 | japan                            | shikoku                                 | sPp     |
| siberian_federal_district               | russia                           | siberian_federal_district               | sPp     |
| sichuan                                 | china                            | sichuan                                 | sPp     |
| sicilia                                 | italy                            | sicilia                                 | sPp     |
| sierra                                  | california                       | sierra                                  | sPp     |
| sikkim                                  | india                            | sikkim                                  | sPp     |
| singapore                               | asia                             | singapore                               | sPp     |
| sint_maarten                            | central-america                  | sint_maarten                            | sPp     |
| siskiyou                                | california                       | siskiyou                                | sPp     |
| slaskie                                 | poland                           | slaskie                                 | sPp     |
| slovakia                                | europe                           | slovakia                                | sPp     |
| smolensk_oblast                         | central_federal_district         | smolensk_oblast                         | sPp     |
| sogn_og_fjordane                        | norway                           | sogn_og_fjordane                        | sPp     |
| solano                                  | california                       | solano                                  | sPp     |
| solomon_islands                         | oceania                          | solomon_islands                         | sPp     |
| somme                                   | picardie                         | somme                                   | sPp     |
| sonoma                                  | california                       | sonoma                                  | sPp     |
| south                                   | brazil                           | south                                   | sPp     |
| south-america                           |                                  | south-america                           | sPp     |
| south_africa                            | africa                           | south_africa                            | sPp     |
| south_australia                         | australia                        | south_australia                         | sPp     |
| south_east                              | england                          | south_east                              | sPp     |
| south_georgia_and_south_sandwich        | south-america                    | south_georgia_and_south_sandwich        | sPp     |
| south_kalimantan                        | indonesia                        | south_kalimantan                        | sPp     |
| south_sudan                             | africa                           | south_sudan                             | sPp     |
| south_sulawesi                          | indonesia                        | south_sulawesi                          | sPp     |
| south_sumatra                           | indonesia                        | south_sumatra                           | sPp     |
| south_west                              | england                          | south_west                              | sPp     |
| southeast                               | brazil                           | southeast                               | sPp     |
| southeast_sulawesi                      | indonesia                        | southeast_sulawesi                      | sPp     |
| southern_federal_district               | russia                           | southern_federal_district               | sPp     |
| southwestern_ontario                    | ontario                          | southwestern_ontario                    | sPp     |
| spain                                   | europe                           | spain                                   | sPp     |
| stanislaus                              | california                       | stanislaus                              | sPp     |
| stavropol_krai                          | north_caucasian_federal_district | stavropol_krai                          | sPp     |
| steiermark                              | austria                          | steiermark                              | sPp     |
| stredocesky                             | czech_republic                   | stredocesky                             | sPp     |
| sudan                                   | africa                           | sudan                                   | sPp     |
| sumy_oblast                             | ukraine                          | sumy_oblast                             | sPp     |
| suriname                                | south-america                    | suriname                                | sPp     |
| sutter                                  | california                       | sutter                                  | sPp     |
| svalbard                                | norway                           | svalbard                                | sPp     |
| sverdlovsk_oblast                       | ural_federal_district            | sverdlovsk_oblast                       | sPp     |
| swaziland                               | africa                           | swaziland                               | sPp     |
| swietokrzyskie                          | poland                           | swietokrzyskie                          | sPp     |
| taaf                                    | south-america                    | taaf                                    | sPp     |
| tambov_oblast                           | central_federal_district         | tambov_oblast                           | sPp     |
| tamil_nadu                              | india                            | tamil_nadu                              | sPp     |
| tarn                                    | midi_pyrenees                    | tarn                                    | sPp     |
| tarn_et_garonne                         | midi_pyrenees                    | tarn_et_garonne                         | sPp     |
| tasmania                                | australia                        | tasmania                                | sPp     |
| tatarstan_republic                      | volga_federal_district           | tatarstan_republic                      | sPp     |
| tehama                                  | california                       | tehama                                  | sPp     |
| telangana                               | india                            | telangana                               | sPp     |
| telemark                                | norway                           | telemark                                | sPp     |
| ternopil_oblast                         | ukraine                          | ternopil_oblast                         | sPp     |
| territoire_de_belfort                   | franche_comte                    | territoire_de_belfort                   | sPp     |
| tianjin                                 | china                            | tianjin                                 | sPp     |
| tibet                                   | china                            | tibet                                   | sPp     |
| tierra_del_fuego                        | argentina                        | tierra_del_fuego                        | sPp     |
| tirol                                   | austria                          | tirol                                   | sPp     |
| tocantins                               | north                            | tocantins                               | sPp     |
| togo                                    | africa                           | togo                                    | sPp     |
| tohoku                                  | japan                            | tohoku                                  | sPp     |
| tokelau                                 | south-america                    | tokelau                                 | sPp     |
| tomsk_oblast                            | siberian_federal_district        | tomsk_oblast                            | sPp     |
| tonga                                   | south-america                    | tonga                                   | sPp     |
| toscana                                 | italy                            | toscana                                 | sPp     |
| trenciansky                             | slovakia                         | trenciansky                             | sPp     |
| trentino_alto_adige                     | italy                            | trentino_alto_adige                     | sPp     |
| trinidad_and_tobago                     | south-america                    | trinidad_and_tobago                     | sPp     |
| trinity                                 | california                       | trinity                                 | sPp     |
| tripura                                 | india                            | tripura                                 | sPp     |
| trnavsky                                | slovakia                         | trnavsky                                | sPp     |
| troendelag                              | norway                           | troendelag                              | sPp     |
| troms                                   | norway                           | troms                                   | sPp     |
| tucuman                                 | argentina                        | tucuman                                 | sPp     |
| tula_oblast                             | central_federal_district         | tula_oblast                             | sPp     |
| tulare                                  | california                       | tulare                                  | sPp     |
| tunisia                                 | africa                           | tunisia                                 | sPp     |
| tuolumne                                | california                       | tuolumne                                | sPp     |
| turks_and_caicos_islands                | central-america                  | turks_and_caicos_islands                | sPp     |
| tuva_republic                           | siberian_federal_district        | tuva_republic                           | sPp     |
| tuvalu                                  | oceania                          | tuvalu                                  | sPp     |
| tver_oblast                             | central_federal_district         | tver_oblast                             | sPp     |
| tyumen_oblast                           | ural_federal_district            | tyumen_oblast                           | sPp     |
| udmurt_republic                         | volga_federal_district           | udmurt_republic                         | sPp     |
| uganda                                  | africa                           | uganda                                  | sPp     |
| ukraine                                 | europe                           | ukraine                                 | sPp     |
| ulyanovsk_oblast                        | volga_federal_district           | ulyanovsk_oblast                        | sPp     |
| umbria                                  | italy                            | umbria                                  | sPp     |
| united_arab_emirates                    | asia                             | united_arab_emirates                    | sPp     |
| united_kingdom                          | europe                           | united_kingdom                          | sPp     |
| united_states_virgin_islands            | central-america                  | united_states_virgin_islands            | sPp     |
| ural_federal_district                   | russia                           | ural_federal_district                   | sPp     |
| us-west                                 | north-america                    | us-west                                 | sPp     |
| usa_virgin_islands                      | central-america                  | usa_virgin_islands                      | sPp     |
| ustecky                                 | czech_republic                   | ustecky                                 | sPp     |
| utrecht                                 | netherlands                      | utrecht                                 | sPp     |
| uttar_pradesh                           | india                            | uttar_pradesh                           | sPp     |
| uttarakhand                             | india                            | uttarakhand                             | sPp     |
| val_d_oise                              | ile_de_france                    | val_d_oise                              | sPp     |
| val_de_marne                            | ile_de_france                    | val_de_marne                            | sPp     |
| valle_aosta                             | italy                            | valle_aosta                             | sPp     |
| vanuatu                                 | oceania                          | vanuatu                                 | sPp     |
| var                                     | provence_alpes_cote_d_azur       | var                                     | sPp     |
| vaucluse                                | provence_alpes_cote_d_azur       | vaucluse                                | sPp     |
| vendee                                  | pays_de_la_loire                 | vendee                                  | sPp     |
| veneto                                  | italy                            | veneto                                  | sPp     |
| venezuela                               | south-america                    | venezuela                               | sPp     |
| ventura                                 | california                       | ventura                                 | sPp     |
| vest-agder                              | norway                           | vest-agder                              | sPp     |
| vestfold                                | norway                           | vestfold                                | sPp     |
| victoria                                | australia                        | victoria                                | sPp     |
| vienne                                  | poitou_charentes                 | vienne                                  | sPp     |
| vinnytsia_oblast                        | ukraine                          | vinnytsia_oblast                        | sPp     |
| vladimir_oblast                         | central_federal_district         | vladimir_oblast                         | sPp     |
| volga_federal_district                  | russia                           | volga_federal_district                  | sPp     |
| volgograd_oblast                        | southern_federal_district        | volgograd_oblast                        | sPp     |
| vologda_oblast                          | northwestern_federal_district    | vologda_oblast                          | sPp     |
| volyn_oblast                            | ukraine                          | volyn_oblast                            | sPp     |
| vorarlberg                              | austria                          | vorarlberg                              | sPp     |
| voronezh_oblast                         | central_federal_district         | voronezh_oblast                         | sPp     |
| vosges                                  | lorraine                         | vosges                                  | sPp     |
| vysocina                                | czech_republic                   | vysocina                                | sPp     |
| wallis_et_futuna                        | south-america                    | wallis_et_futuna                        | sPp     |
| wallonia_french_community               | belgium                          | wallonia_french_community               | sPp     |
| wallonia_german_community               | belgium                          | wallonia_german_community               | sPp     |
| warminsko_mazurskie                     | poland                           | warminsko_mazurskie                     | sPp     |
| west_bengal                             | india                            | west_bengal                             | sPp     |
| west_java                               | indonesia                        | west_java                               | sPp     |
| west_kalimantan                         | indonesia                        | west_kalimantan                         | sPp     |
| west_midlands                           | england                          | west_midlands                           | sPp     |
| west_nusa_tenggara                      | indonesia                        | west_nusa_tenggara                      | sPp     |
| west_papua                              | indonesia                        | west_papua                              | sPp     |
| west_sulawesi                           | indonesia                        | west_sulawesi                           | sPp     |
| west_sumatra                            | indonesia                        | west_sumatra                            | sPp     |
| western_australia                       | australia                        | western_australia                       | sPp     |
| western_sahara                          | africa                           | western_sahara                          | sPp     |
| wielkopolskie                           | poland                           | wielkopolskie                           | sPp     |
| wien                                    | austria                          | wien                                    | sPp     |
| xinjiang                                | china                            | xinjiang                                | sPp     |
| yamalo_nenets_autonomous_okrug          | ural_federal_district            | yamalo_nenets_autonomous_okrug          | sPp     |
| yaroslavl_oblast                        | central_federal_district         | yaroslavl_oblast                        | sPp     |
| yogyakarta                              | indonesia                        | yogyakarta                              | sPp     |
| yolo                                    | california                       | yolo                                    | sPp     |
| yonne                                   | bourgogne                        | yonne                                   | sPp     |
| yorkshire_and_the_humber                | england                          | yorkshire_and_the_humber                | sPp     |
| yuba                                    | california                       | yuba                                    | sPp     |
| yukon                                   | canada                           | yukon                                   | sPp     |
| yunnan                                  | china                            | yunnan                                  | sPp     |
| yvelines                                | ile_de_france                    | yvelines                                | sPp     |
| zabaykalsky_krai                        | siberian_federal_district        | zabaykalsky_krai                        | sPp     |
| zachodniopomorskie                      | poland                           | zachodniopomorskie                      | sPp     |
| zakarpattia_oblast                      | ukraine                          | zakarpattia_oblast                      | sPp     |
| zambia                                  | africa                           | zambia                                  | sPp     |
| zaporizhia_oblast                       | ukraine                          | zaporizhia_oblast                       | sPp     |
| zeeland                                 | netherlands                      | zeeland                                 | sPp     |
| zhejiang                                | china                            | zhejiang                                | sPp     |
| zhytomyr_oblast                         | ukraine                          | zhytomyr_oblast                         | sPp     |
| zilinsky                                | slovakia                         | zilinsky                                | sPp     |
| zimbabwe                                | africa                           | zimbabwe                                | sPp     |
| zlinsky                                 | czech_republic                   | zlinsky                                 | sPp     |
| zuid_holland                            | netherlands                      | zuid_holland                            | sPp     |
Total elements: 833

## List of elements from bbbike.org
|    SHORTNAME     | IS IN |    LONG NAME     | FORMATS |
|------------------|-------|------------------|---------|
| Aachen           |       | Aachen           | PGS     |
| Aarhus           |       | Aarhus           | PGS     |
| Adelaide         |       | Adelaide         | PGS     |
| Albuquerque      |       | Albuquerque      | PGS     |
| Alexandria       |       | Alexandria       | PGS     |
| Amsterdam        |       | Amsterdam        | PGS     |
| Antwerpen        |       | Antwerpen        | PGS     |
| Arnhem           |       | Arnhem           | PGS     |
| Auckland         |       | Auckland         | PGS     |
| Augsburg         |       | Augsburg         | PGS     |
| Austin           |       | Austin           | PGS     |
| Baghdad          |       | Baghdad          | PGS     |
| Baku             |       | Baku             | PGS     |
| Balaton          |       | Balaton          | PGS     |
| Bamberg          |       | Bamberg          | PGS     |
| Bangkok          |       | Bangkok          | PGS     |
| Barcelona        |       | Barcelona        | PGS     |
| Basel            |       | Basel            | PGS     |
| Beijing          |       | Beijing          | PGS     |
| Beirut           |       | Beirut           | PGS     |
| Berkeley         |       | Berkeley         | PGS     |
| Berlin           |       | Berlin           | PGS     |
| Bern             |       | Bern             | PGS     |
| Bielefeld        |       | Bielefeld        | PGS     |
| Birmingham       |       | Birmingham       | PGS     |
| Bochum           |       | Bochum           | PGS     |
| Bogota           |       | Bogota           | PGS     |
| Bombay           |       | Bombay           | PGS     |
| Bonn             |       | Bonn             | PGS     |
| Bordeaux         |       | Bordeaux         | PGS     |
| Boulder          |       | Boulder          | PGS     |
| BrandenburgHavel |       | BrandenburgHavel | PGS     |
| Braunschweig     |       | Braunschweig     | PGS     |
| Bremen           |       | Bremen           | PGS     |
| Bremerhaven      |       | Bremerhaven      | PGS     |
| Brisbane         |       | Brisbane         | PGS     |
| Bristol          |       | Bristol          | PGS     |
| Brno             |       | Brno             | PGS     |
| Bruegge          |       | Bruegge          | PGS     |
| Bruessel         |       | Bruessel         | PGS     |
| Budapest         |       | Budapest         | PGS     |
| BuenosAires      |       | BuenosAires      | PGS     |
| Cairo            |       | Cairo            | PGS     |
| Calgary          |       | Calgary          | PGS     |
| Cambridge        |       | Cambridge        | PGS     |
| CambridgeMa      |       | CambridgeMa      | PGS     |
| Canberra         |       | Canberra         | PGS     |
| CapeTown         |       | CapeTown         | PGS     |
| Chemnitz         |       | Chemnitz         | PGS     |
| Chicago          |       | Chicago          | PGS     |
| ClermontFerrand  |       | ClermontFerrand  | PGS     |
| Colmar           |       | Colmar           | PGS     |
| Copenhagen       |       | Copenhagen       | PGS     |
| Cork             |       | Cork             | PGS     |
| Corsica          |       | Corsica          | PGS     |
| Corvallis        |       | Corvallis        | PGS     |
| Cottbus          |       | Cottbus          | PGS     |
| Cracow           |       | Cracow           | PGS     |
| CraterLake       |       | CraterLake       | PGS     |
| Curitiba         |       | Curitiba         | PGS     |
| Cusco            |       | Cusco            | PGS     |
| Dallas           |       | Dallas           | PGS     |
| Darmstadt        |       | Darmstadt        | PGS     |
| Davis            |       | Davis            | PGS     |
| DenHaag          |       | DenHaag          | PGS     |
| Denver           |       | Denver           | PGS     |
| Dessau           |       | Dessau           | PGS     |
| Dortmund         |       | Dortmund         | PGS     |
| Dresden          |       | Dresden          | PGS     |
| Dublin           |       | Dublin           | PGS     |
| Duesseldorf      |       | Duesseldorf      | PGS     |
| Duisburg         |       | Duisburg         | PGS     |
| Edinburgh        |       | Edinburgh        | PGS     |
| Eindhoven        |       | Eindhoven        | PGS     |
| Emden            |       | Emden            | PGS     |
| Erfurt           |       | Erfurt           | PGS     |
| Erlangen         |       | Erlangen         | PGS     |
| Eugene           |       | Eugene           | PGS     |
| Flensburg        |       | Flensburg        | PGS     |
| FortCollins      |       | FortCollins      | PGS     |
| Frankfurt        |       | Frankfurt        | PGS     |
| FrankfurtOder    |       | FrankfurtOder    | PGS     |
| Freiburg         |       | Freiburg         | PGS     |
| Gdansk           |       | Gdansk           | PGS     |
| Genf             |       | Genf             | PGS     |
| Gent             |       | Gent             | PGS     |
| Gera             |       | Gera             | PGS     |
| Glasgow          |       | Glasgow          | PGS     |
| Gliwice          |       | Gliwice          | PGS     |
| Goerlitz         |       | Goerlitz         | PGS     |
| Goeteborg        |       | Goeteborg        | PGS     |
| Goettingen       |       | Goettingen       | PGS     |
| Graz             |       | Graz             | PGS     |
| Groningen        |       | Groningen        | PGS     |
| Halifax          |       | Halifax          | PGS     |
| Halle            |       | Halle            | PGS     |
| Hamburg          |       | Hamburg          | PGS     |
| Hamm             |       | Hamm             | PGS     |
| Hannover         |       | Hannover         | PGS     |
| Heilbronn        |       | Heilbronn        | PGS     |
| Helsinki         |       | Helsinki         | PGS     |
| Hertogenbosch    |       | Hertogenbosch    | PGS     |
| Huntsville       |       | Huntsville       | PGS     |
| Innsbruck        |       | Innsbruck        | PGS     |
| Istanbul         |       | Istanbul         | PGS     |
| Jena             |       | Jena             | PGS     |
| Jerusalem        |       | Jerusalem        | PGS     |
| Johannesburg     |       | Johannesburg     | PGS     |
| Kaiserslautern   |       | Kaiserslautern   | PGS     |
| Karlsruhe        |       | Karlsruhe        | PGS     |
| Kassel           |       | Kassel           | PGS     |
| Katowice         |       | Katowice         | PGS     |
| Kaunas           |       | Kaunas           | PGS     |
| Kiel             |       | Kiel             | PGS     |
| Kiew             |       | Kiew             | PGS     |
| Koblenz          |       | Koblenz          | PGS     |
| Koeln            |       | Koeln            | PGS     |
| Konstanz         |       | Konstanz         | PGS     |
| LaPaz            |       | LaPaz            | PGS     |
| LaPlata          |       | LaPlata          | PGS     |
| LakeGarda        |       | LakeGarda        | PGS     |
| Lausanne         |       | Lausanne         | PGS     |
| Leeds            |       | Leeds            | PGS     |
| Leipzig          |       | Leipzig          | PGS     |
| Lima             |       | Lima             | PGS     |
| Linz             |       | Linz             | PGS     |
| Lisbon           |       | Lisbon           | PGS     |
| Liverpool        |       | Liverpool        | PGS     |
| Ljubljana        |       | Ljubljana        | PGS     |
| Lodz             |       | Lodz             | PGS     |
| London           |       | London           | PGS     |
| Luebeck          |       | Luebeck          | PGS     |
| Luxemburg        |       | Luxemburg        | PGS     |
| Lyon             |       | Lyon             | PGS     |
| Maastricht       |       | Maastricht       | PGS     |
| Madison          |       | Madison          | PGS     |
| Madrid           |       | Madrid           | PGS     |
| Magdeburg        |       | Magdeburg        | PGS     |
| Mainz            |       | Mainz            | PGS     |
| Malmoe           |       | Malmoe           | PGS     |
| Manchester       |       | Manchester       | PGS     |
| Mannheim         |       | Mannheim         | PGS     |
| Marseille        |       | Marseille        | PGS     |
| Melbourne        |       | Melbourne        | PGS     |
| Memphis          |       | Memphis          | PGS     |
| MexicoCity       |       | MexicoCity       | PGS     |
| Miami            |       | Miami            | PGS     |
| Moenchengladbach |       | Moenchengladbach | PGS     |
| Montevideo       |       | Montevideo       | PGS     |
| Montpellier      |       | Montpellier      | PGS     |
| Montreal         |       | Montreal         | PGS     |
| Moscow           |       | Moscow           | PGS     |
| Muenchen         |       | Muenchen         | PGS     |
| Muenster         |       | Muenster         | PGS     |
| NewDelhi         |       | NewDelhi         | PGS     |
| NewOrleans       |       | NewOrleans       | PGS     |
| NewYork          |       | NewYork          | PGS     |
| Nuernberg        |       | Nuernberg        | PGS     |
| Oldenburg        |       | Oldenburg        | PGS     |
| Oranienburg      |       | Oranienburg      | PGS     |
| Orlando          |       | Orlando          | PGS     |
| Oslo             |       | Oslo             | PGS     |
| Osnabrueck       |       | Osnabrueck       | PGS     |
| Ostrava          |       | Ostrava          | PGS     |
| Ottawa           |       | Ottawa           | PGS     |
| Paderborn        |       | Paderborn        | PGS     |
| Palma            |       | Palma            | PGS     |
| PaloAlto         |       | PaloAlto         | PGS     |
| Paris            |       | Paris            | PGS     |
| Perth            |       | Perth            | PGS     |
| Philadelphia     |       | Philadelphia     | PGS     |
| PhnomPenh        |       | PhnomPenh        | PGS     |
| Portland         |       | Portland         | PGS     |
| PortlandME       |       | PortlandME       | PGS     |
| Porto            |       | Porto            | PGS     |
| PortoAlegre      |       | PortoAlegre      | PGS     |
| Potsdam          |       | Potsdam          | PGS     |
| Poznan           |       | Poznan           | PGS     |
| Prag             |       | Prag             | PGS     |
| Providence       |       | Providence       | PGS     |
| Regensburg       |       | Regensburg       | PGS     |
| Riga             |       | Riga             | PGS     |
| RiodeJaneiro     |       | RiodeJaneiro     | PGS     |
| Rostock          |       | Rostock          | PGS     |
| Rotterdam        |       | Rotterdam        | PGS     |
| Ruegen           |       | Ruegen           | PGS     |
| Saarbruecken     |       | Saarbruecken     | PGS     |
| Sacramento       |       | Sacramento       | PGS     |
| Saigon           |       | Saigon           | PGS     |
| Salzburg         |       | Salzburg         | PGS     |
| SanFrancisco     |       | SanFrancisco     | PGS     |
| SanJose          |       | SanJose          | PGS     |
| SanktPetersburg  |       | SanktPetersburg  | PGS     |
| SantaBarbara     |       | SantaBarbara     | PGS     |
| SantaCruz        |       | SantaCruz        | PGS     |
| Santiago         |       | Santiago         | PGS     |
| Sarajewo         |       | Sarajewo         | PGS     |
| Schwerin         |       | Schwerin         | PGS     |
| Seattle          |       | Seattle          | PGS     |
| Seoul            |       | Seoul            | PGS     |
| Sheffield        |       | Sheffield        | PGS     |
| Singapore        |       | Singapore        | PGS     |
| Sofia            |       | Sofia            | PGS     |
| Stockholm        |       | Stockholm        | PGS     |
| Stockton         |       | Stockton         | PGS     |
| Strassburg       |       | Strassburg       | PGS     |
| Stuttgart        |       | Stuttgart        | PGS     |
| Sucre            |       | Sucre            | PGS     |
| Sydney           |       | Sydney           | PGS     |
| Szczecin         |       | Szczecin         | PGS     |
| Tallinn          |       | Tallinn          | PGS     |
| Tehran           |       | Tehran           | PGS     |
| Tilburg          |       | Tilburg          | PGS     |
| Tokyo            |       | Tokyo            | PGS     |
| Toronto          |       | Toronto          | PGS     |
| Toulouse         |       | Toulouse         | PGS     |
| Trondheim        |       | Trondheim        | PGS     |
| Tucson           |       | Tucson           | PGS     |
| Turin            |       | Turin            | PGS     |
| UlanBator        |       | UlanBator        | PGS     |
| Ulm              |       | Ulm              | PGS     |
| Usedom           |       | Usedom           | PGS     |
| Utrecht          |       | Utrecht          | PGS     |
| Vancouver        |       | Vancouver        | PGS     |
| Victoria         |       | Victoria         | PGS     |
| WarenMueritz     |       | WarenMueritz     | PGS     |
| Warsaw           |       | Warsaw           | PGS     |
| WashingtonDC     |       | WashingtonDC     | PGS     |
| Waterloo         |       | Waterloo         | PGS     |
| Wien             |       | Wien             | PGS     |
| Wroclaw          |       | Wroclaw          | PGS     |
| Wuerzburg        |       | Wuerzburg        | PGS     |
| Wuppertal        |       | Wuppertal        | PGS     |
| Zagreb           |       | Zagreb           | PGS     |
| Zuerich          |       | Zuerich          | PGS     |
Total elements: 235

