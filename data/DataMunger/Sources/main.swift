import SwiftyJSON
import Foundation

let inputFolder = "/Users/McJones/Downloads/"
let outputFolder = "/Users/McJones/Development/govhack2017/data/"

// basic function matches keywords to death types
// not perfect but avoids us having to make ~1500 death cards
func deathTypes(death:String) -> [String]
{
    var deaths : Set<String> = []
    
    let types = ["DROWN":"dc_drowning",
                 "ACCIDENT":"dc_misc",
                 "BOWEL":"dc_stomach",
                 "SUICIDE":"dc_suicide",
                 "NATURAL":"dc_misc",
                 "ACID":"dc_acid",
                 "OLD AGE":"dc_age",
                 "ALCOHOL":"dc_alcohol",
                 "BRAIN":"dc_brain",
                 "BRIDGE":"dc_brain",
                 "BULL":"dc_bull",
                 "BULLOCK":"dc_bull",
                 "BURN":"dc_burning",
                 "CRUSH":"dc_crush",
                 "EXPLO":"dc_explosion",
                 "EXPO":"dc_explosion",
                 "FALL":"dc_falling",
                 "GUN":"dc_gunshot",
                 "SHOT":"dc_gunshot",
                 "HEART":"dc_heart",
                 "CARDIAC":"dc_heart",
                 "HORSE":"dc_horse",
                 "DISEASE":"dc_illness",
                 "POISON":"dc_poison",
                 "SCALD":"dc_scalding",
                 "LUNG":"dc_lungs",
                 "INHAL":"dc_inhalation",
                 "MEDIC":"dc_medic",
                 "MINE":"dc_mining",
                 "MINI":"dc_mining",
                 "SHIP":"dc_shipwreck",
                 "SHOCK":"dc_shock",
                 "SNAKE":"dc_snakebite",
                 "STARV":"dc_starvation",
                 "STOMACH":"dc_stomach",
                 "TRAIN":"dc_train",
                 "RAILWAY":"dc_train",
                 "TREE":"dc_tree",
                 "VISIT":"dc_visitation",
                 "WAGON":"dc_wagon",
                 "WAGGON":"dc_wagon",
                 "WHEEL":"dc_wheel",
                 "MURDER":"dc_stabbing",
                 "DIARR":"dc_stomach",
                 "EPILEP":"dc_brain",
                 "APOPLEXY":"dc_brain",
                 "SELF":"dc_suicide",
                 "CHILD":"dc_childbirth",
                 "DRINK":"dc_alcohol",
                 "DIP":"dc_illness",
                 "CONVULS":"dc_brain",
                 "DRIV":"dc_wheel",
                 "PNEUMONIA":"dc_lungs",
                 "BRONCH":"dc_lungs",
                 "SUFFO":"dc_lungs",
                 "HOMICIDE":"dc_stabbing",
                 "SENIL":"dc_age",
                 "DROPSY":"dc_illness",
                 "NOURISHMENT":"dc_stomach",
                 "RIVER":"dc_drowning",
                 "INFLUENZA":"dc_illness",
                 "CANCER":"dc_illness",
                 "SYNCOPE":"dc_heart",
                 "MENTAL":"dc_brain",
                 "CEREBRAL":"dc_brain",
                 "HAEMORRHAGE":"dc_brain",
                 "ASTH":"dc_lungs",
                 "BLOOD":"dc_heart",
                 "RUPTURE":"dc_blood",
                 "MANSLAUGHTER":"dc_stabbing",
                 "LAUDANUM":"dc_poison",
                 "BIRTH":"dc_childbirth",
                 "MISADVENTURE":"dc_misadventure",
                 "OVERDOSE":"dc_poison",
                 "DECAY":"dc_age",
                 "FOUND":"dc_unexplained",
                 "SKULL":"dc_brain",
                 "NUTRITION":"dc_starvation",
                 "EXHAUST":"dc_exposure",
                 "FOOD":"dc_starvation",
                 "PERITONITIS":"dc_stomach",
                 "DYSENT":"dc_stomach",
                 "TETANUS":"dc_illness",
                 " OWN ":"dc_suicide",
                 "NOT KNOWN":"dc_unexplained",
                 "UNKNOW":"dc_unexplained",
                 "NOXI":"dc_poison",
                 "HEAD":"dc_brain",
                 "CHLOROFORM":"dc_poison",
                 "ASPHYXIA":"dc_lungs",
                 "COUGH":"dc_lungs",
                 "DOSE":"dc_poison",
                 "CONSUMPTION":"dc_starvation",
                 "PULMONARY":"dc_heart",
                 "ATROPHY":"dc_exposure",
                 "PERICARDIUM":"dc_heart",
                 "NO EVIDENCE":"dc_unexplained",
                 "SEPTICAEMIA":"dc_poison",
                 "RUN OVER":"dc_wheel",
                 "INTESTINE":"dc_stomach",
                 "ABSCESS":"dc_stomach",
                 "KIDNEY":"dc_stomach",
                 "DISLOCATION":"dc_falling",
                 "PLEURISY":"dc_lungs",
                 "EATING":"dc_stomach",
                 "VERTIGO":"dc_falling",
                 "SUNSTROKE":"dc_illness",
                 "MENINGITIS":"dc_illness"]
    
    for (deathKey,deathType) in types
    {
        if death.uppercased().contains(deathKey)
        {
            deaths.insert(deathType)
        }
    }
    
    // they didn't get picked up yet
    // then it means they are one of those "hey it's murder!" leftover deaths
    if deaths.count == 0
    {
        deaths.insert("dc_stabbing")
    }
    
    return Array(deaths)
}

func jsonThatShit(_ jsonName : String) -> JSON
{
    let jsonURL = URL(fileURLWithPath: "\(inputFolder)\(jsonName).json")
    let jsonData = try! Data(contentsOf: jsonURL)
    return JSON(data: jsonData)
}
func mungeJSON(_ jsonName : String, requiredKeys : [LincKey], optionalKeys : [LincKey]?)
{
    for (_,json) in jsonThatShit(jsonName)
    {
        if let name = json[LincKey.name.rawValue].string
        {
            if var newData = data[name]
            {
                for required in requiredKeys
                {
                    var newKey = required
                    guard let reqVal = json[required.rawValue].string else
                    {
                        break
                    }
                    
                    // if it is one of the generic types
                    if required == .permalink || required == .year
                    {
                        newKey = LincKey(rawValue: jsonName + "_" + required.rawValue)!
                    }
                    
                    newData[newKey] = reqVal
                }
                if let optionalKeys = optionalKeys
                {
                    for optional in optionalKeys
                    {
                        var newKey = optional
                        if let optVal = json[optional.rawValue].string
                        {
                            // if it is one of the generic types
                            if optional == .permalink || optional == .year
                            {
                                newKey = LincKey(rawValue: jsonName + "_" + optional.rawValue)!
                            }
                            
                            newData[newKey] = optVal
                        }
                    }
                }
                data[name] = newData
            }
        }
    }
}

func wipeBirthData(person : [LincKey : String])->[LincKey : String]
{
    var peep = person
    
    peep[.birthDate] = nil
    peep[.birthPlace] = nil
    peep[.birthMother] = nil
    peep[.birthFather] = nil
    peep[.birthsYear] = nil
    peep[.birthsPermalink] = nil
    
    return peep
}
func wipeConvictData(person : [LincKey : String])->[LincKey : String]
{
    var peep = person
    
    peep[.convictDate] = nil
    peep[.convictPort] = nil
    peep[.convictShip] = nil
    peep[.convictYear] = nil
    peep[.convictPermalink] = nil
    
    return peep
}
func wipeImmigrationData(person : [LincKey : String])->[LincKey : String]
{
    var peep = person
    
    peep[.immigrationDate] = nil
    peep[.immigrationOrigin] = nil
    peep[.immigrationYear] = nil
    peep[.immigrationPermalink] = nil
    
    return peep
}


// holds the required keys for the json
// made because I constantly make typos
enum LincKey : String
{
    // used all over the place
    case name = "NAME_FULL_DISPLAY"
    case year = "YEAR"
    case permalink = "PERMA_LINK"
    
    // inquest
    case deathDate = "DEATH_DATE"
    case deathVerdict = "VERDICT"
    // generated
    case deathYear = "inquests_YEAR"
    case deathPermalink = "inquests_PERMA_LINK"
    
    // spawn events
    case birthDate = "BIRTH_DATE"
    case birthPlace = "REG_PLACE"
    case birthMother = "MOTHER"
    case birthFather = "FATHER"
    // generated
    case birthsYear = "births_YEAR"
    case birthsPermalink = "births_PERMA_LINK"
    
    case convictDate = "DEPARTURE_DATE"
    case convictPort = "DEPARTURE_PORT"
    case convictShip = "SHIP"
    // generated
    case convictYear = "convicts_YEAR"
    case convictPermalink = "convicts_PERMA_LINK"
    
    case immigrationDate = "DOC_DATE"
    case immigrationOrigin = "SHIP_NATIVE_PLACE"
    // generated
    case immigrationYear = "immigration_YEAR"
    case immigrationPermalink = "immigration_PERMA_LINK"
    
    // life events
    case marriageDate = "MARRIAGE_DATE"
    case marriageAge = "MARRIAGE_AGE"
    case marriageSpouse = "SPOUSE"
    // generated
    case marriagesYear = "marriages_YEAR"
    case marriagesPermalink = "marriages_PERMA_LINK"
    
    case bankruptDate = "BANK_DATE"
    // generated
    case bankruptYear = "bankruptcy_YEAR"
    case bankruptPermalink = "bankruptcy_PERMA_LINK"
    
    case courtDate = "TRIAL_DATE"
    case courtLocation = "TRIAL_PLACE"
    case courtOffense = "OFFENSE"
    // generated
    case courtYear = "court_YEAR"
    case courtPermalink = "court_PERMA_LINK"
    
    case hospitalDate = "ADMISS_DATE"
    case hospitalRemark = "REMARKS"
    case hospitalPlace = "PROPERTY"
    // generated
    case hospitalYear = "health--welfare_YEAR"
    case hospitalPermalink = "health--welfare_PERMA_LINK"
    
    case censusKids = "UNDER14"
    case censusPlace = "CENSUS_DISTRICT"
    // generated
    case censusYear = "census_YEAR"
    case censusPermalink = "census_PERMA_LINK"
}

var data : [String:[LincKey:String]] = [:]

var deaths : Set<String> = []

for (key, json) in jsonThatShit("inquests")
{
    if let name = json["NAME"].string,
       let deathDate = json[LincKey.deathDate.rawValue].string,
       let deathReason = json[LincKey.deathVerdict.rawValue].string
    {
        var death : [LincKey:String] = [.deathDate : deathDate, .deathVerdict : deathReason]
        if let deathYear = json[LincKey.year.rawValue].string
        {
            death[.deathYear] = deathYear
        }
        if let deathPerma = json[LincKey.permalink.rawValue].string
        {
            death[.deathPermalink] = deathPerma
        }
        data[name] = death
        
        deaths.insert(deathReason)
    }
}

mungeJSON("convicts", requiredKeys: [.convictDate,.convictPort], optionalKeys: [.convictShip, .year, .permalink])

mungeJSON("births", requiredKeys: [.birthDate,.birthPlace], optionalKeys: [.year, .permalink, .birthFather, .birthMother])

mungeJSON("immigration", requiredKeys: [.immigrationDate], optionalKeys: [.immigrationOrigin,.year,.permalink])

mungeJSON("marriages", requiredKeys: [.marriageDate], optionalKeys: [.marriageAge,.year,.permalink,.marriageSpouse])

mungeJSON("bankruptcy", requiredKeys: [.bankruptDate], optionalKeys: [.year,.permalink])

mungeJSON("court", requiredKeys: [.courtOffense,.courtDate,.courtLocation], optionalKeys: [.year,.permalink])

mungeJSON("health--welfare", requiredKeys: [.hospitalPlace,.hospitalRemark], optionalKeys: [.year,.permalink])

mungeJSON("census", requiredKeys: [.censusKids,.censusPlace], optionalKeys: [.year,.permalink])

// "cleaning" up the data
// currently everyone needs a spawning event and only one
// dropping them in priority of convict, immigrant, birth
for person in data
{
    var peep : [LincKey : String]? = nil
    
    if person.value[.immigrationDate] != nil
    {
        peep = wipeBirthData(person: wipeConvictData(person: person.value))
    }
    else if person.value[.convictDate] != nil
    {
        peep = wipeImmigrationData(person: wipeBirthData(person: person.value))
    }
    else if person.value[.birthDate] != nil
    {
        peep = wipeImmigrationData(person: wipeConvictData(person: person.value))
    }
    
    data[person.key] = peep
}

// debugging output
print("Total people: \(data.count)")
print("Tassie born: \(data.filter({ $0.value[.birthDate] != nil }).count)")
print("Immigrant: \(data.filter({ $0.value[.immigrationDate] != nil }).count)")
print("Convicts: \(data.filter({ $0.value[.convictDate] != nil }).count)")
print("Hitched people: \(data.filter({ $0.value[.marriageDate] != nil }).count)")
print("Bankrupt people: \(data.filter({ $0.value[.bankruptDate] != nil }).count)")
print("Trialed people: \(data.filter({ $0.value[.courtOffense] != nil }).count)")
print("Hospitalised people: \(data.filter({ $0.value[.hospitalRemark] != nil }).count)")
print("Censused people: \(data.filter({ $0.value[.censusPlace] != nil }).count)")
print("Had kids at census time: \(data.filter({ $0.value[.censusKids] == "Yes" }).count)")

print("\nThere are \(deaths.count) deaths")
//print("There are \(deaths.filter({ $0.contains("Accident") }).count) accidental deaths")
//print("There are \(deaths.filter({ $0.contains("Accidentally drowned") }).count) accidental drownings\n")
//deaths.sorted().forEach({ print($0) })

// outputting the data to the form the server needs
var json : [[String:Any]] = []
for (key,value) in data
{
    var newPerson : [String:Any] = [:]
    newPerson["name"] = key
    
    var inquest : [String:Any] = [:]
    
    var birth : [String:String] = [:]
    
    var convict : [String:String] = [:]
    
    var immigration : [String:String] = [:]
    
    var marriage : [String:String] = [:]
    
    var bankruptcy : [String:String] = [:]
    
    var court : [String:String] = [:]
    
    var health : [String:String] = [:]
    
    var census : [String:Any] = [:]
    
    // ugly but gives me full control...
    for (dataKey, dataValue) in value
    {
        switch dataKey
        {
        case .deathDate:
            inquest["death_date"] = dataValue
        case .deathVerdict:
            inquest["death_verdict"] = dataValue
            inquest["death_causes"] = deathTypes(death: dataValue)
        case .deathYear:
            inquest["year"] = dataValue
        case .deathPermalink:
            inquest["permalink"] = dataValue
            
        case .birthDate:
            birth["birth_date"] = dataValue
        case .birthPlace:
            birth["birth_place"] = dataValue
        case .birthMother:
            birth["birth_mother"] = dataValue
        case .birthFather:
            birth["birth_father"] = dataValue
        case .birthsYear:
            birth["year"] = dataValue
        case .birthsPermalink:
            birth["permalink"] = dataValue
        
        case .convictDate:
            convict["departure_date"] = dataValue
        case .convictPort:
            convict["convict_port"] = dataValue
        case .convictShip:
            convict["convict_ship"] = dataValue
        case .convictYear:
            convict["year"] = dataValue
        case .convictPermalink:
            convict["permalink"] = dataValue
            
        case .immigrationDate:
            immigration["immigration_date"] = dataValue
        case .immigrationOrigin:
            immigration["from_country"] = dataValue
        case .immigrationYear:
            immigration["year"] = dataValue
        case .immigrationPermalink:
            immigration["permalink"] = dataValue
            
        case .marriageDate:
            marriage["marriage_date"] = dataValue
        case .marriageAge:
            marriage["marriage_age"] = dataValue
        case .marriageSpouse:
            marriage["spouse_name"] = dataValue
        case .marriagesYear:
            marriage["year"] = dataValue
        case .marriagesPermalink:
            marriage["permalink"] = dataValue
            
        case .bankruptDate:
            bankruptcy["bankrupt_date"] = dataValue
        case .bankruptYear:
            bankruptcy["year"] = dataValue
        case .bankruptPermalink:
            bankruptcy["permalink"] = dataValue
            
        case .courtDate:
            court["trial_date"] = dataValue
        case .courtLocation:
            court["trial_location"] = dataValue
        case .courtOffense:
            court["trial_offence"] = dataValue
        case .courtYear:
            court["year"] = dataValue
        case .courtPermalink:
            court["permalink"] = dataValue
            
        case .hospitalDate:
            health["admission_date"] = dataValue
        case .hospitalRemark:
            health["remarks"] = dataValue
        case .hospitalPlace:
            health["property"] = dataValue
        case .hospitalYear:
            health["year"] = dataValue
        case .hospitalPermalink:
            health["permalink"] = dataValue
            
        case .censusKids:
            census["census_children"] = dataValue == "Yes" ? true : false
        case .censusPlace:
            census["census_place"] = dataValue
        case .censusYear:
            census["year"] = dataValue
            census["census_year"] = dataValue
        case .censusPermalink:
            census["permalink"] = dataValue
            
        default:
            fatalError("This should not happen")
        }
    }
    
    newPerson["inquest"] = inquest
    newPerson["birth"] = birth.count > 0 ? birth : nil
    newPerson["immigration"] = immigration.count > 0 ? immigration : nil
    newPerson["convict"] = convict.count > 0 ? convict : nil
    newPerson["bankruptcy"] = bankruptcy.count > 0 ? bankruptcy : nil
    newPerson["marriage"] = marriage.count > 0 ? marriage : nil
    newPerson["court"] = court.count > 0 ? court : nil
    newPerson["health-welfare"] = health.count > 0 ? health : nil
    newPerson["census"] = census.count > 0 ? census : nil
    
    json.append(newPerson)
}
let jsonData = try! JSONSerialization.data(withJSONObject: json, options: .prettyPrinted)
let outputURL = URL(fileURLWithPath: "\(outputFolder)person.json")
try! jsonData.write(to: outputURL)

for person in json
{
    guard let inquest = person["inquest"] as? [String:Any] else
    {
        fatalError("No inquest found for \(person["name"])")
    }
    guard let deathTypes = inquest["death_causes"] as? [String] else
    {
        fatalError("No death causes found for \(person["name"])")
    }
    guard deathTypes.count > 0 else
    {
        fatalError("\(person["name"]) has no death type: \(inquest["death_verdict"]))")
    }
}

