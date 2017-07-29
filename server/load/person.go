package load

// Person is used to load the people.
type Person struct {
	Name    string `json:"name"`
	Inquest struct {
		DeathDate    string   `json:"death_date"`
		DeathVerdict string   `json:"death_verdict"`
		DeathCauses  []string `json:"death_causes"`
		Year         string   `json:"year"`
		Permalink    string   `json:"permalink"`
	} `json:"inquest"`
	Birth struct {
		BirthDate   string `json:"birth_date"`
		BirthPlace  string `json:"birth_place"`
		BirthMother string `json:"birth_mother"`
		BirthFather string `json:"birth_father"`
		Year        string `json:"year"`
		Permalink   string `json:"permalink"`
	} `json:"birth"`
	Immigration struct {
		ImmigrationDate string `json:"immigration_date"`
		FromCountry     string `json:"from_country"`
		Year            string `json:"year"`
		Permalink       string `json:"permalink"`
	} `json:"immigration"`
	Convict struct {
		DepartureDate string `json:"departure_date"`
		ConvictPort   string `json:"convict_port"`
		ConvictShip   string `json:"convict_ship"`
		Year          string `json:"year"`
		Permalink     string `json:"permalink"`
	} `json:"convict"`
	Bankruptcy struct {
		BankruptDate string `json:"bankrupt_date"`
		Year         string `json:"year"`
		Permalink    string `json:"permalink"`
	} `json:"bankruptcy"`
	Marriage struct {
		MarriageDate  string `json:"marriage_date"`
		SpouseName    string `json:"spouse_name"`
		MarriagePlace string `json:"marriage_place"`
		Year          string `json:"year"`
		Permalink     string `json:"permalink"`
	} `json:"marriage"`
	Court struct {
		TrialDate    string `json:"trial_date"`
		TrialOffence string `json:"trial_offence"`
		Year         string `json:"year"`
		Permalink    string `json:"permalink"`
	} `json:"court"`
	HealthWelfare struct {
		AdmissionDate string `json:"admission_date"`
		Property      string `json:"property"`
		Remarks       string `json:"remarks"`
		Year          string `json:"year"`
		Permalink     string `json:"permalink"`
	} `json:"health-welfare"`
	Census struct {
		CensusYear     string `json:"census_year"`
		CensusPlace    string `json:"census_place"`
		CensusChildren bool   `json:"census_children"`
		Year           string `json:"year"`
		Permalink      string `json:"permalink"`
	} `json:"census"`
}
