package load

// Person is used to load the people.
type Person struct {
	Name    string `json:"name"`
	Inquest struct {
		DeathDate    string   `json:"death_date,omitempty"`
		DeathVerdict string   `json:"death_verdict,omitempty"`
		DeathCauses  []string `json:"death_causes,omitempty"`
		Year         string   `json:"year,omitempty"`
		Permalink    string   `json:"permalink,omitempty"`
	} `json:"inquest,omitempty"`
	Birth struct {
		BirthDate   string `json:"birth_date,omitempty"`
		BirthPlace  string `json:"birth_place,omitempty"`
		BirthMother string `json:"birth_mother,omitempty"`
		BirthFather string `json:"birth_father,omitempty"`
		Year        string `json:"year,omitempty"`
		Permalink   string `json:"permalink,omitempty"`
	} `json:"birth,omitempty"`
	Immigration struct {
		ImmigrationDate string `json:"immigration_date,omitempty"`
		FromCountry     string `json:"from_country,omitempty"`
		Year            string `json:"year,omitempty"`
		Permalink       string `json:"permalink,omitempty"`
	} `json:"immigration,omitempty"`
	Convict struct {
		DepartureDate string `json:"departure_date,omitempty"`
		ConvictPort   string `json:"convict_port,omitempty"`
		ConvictShip   string `json:"convict_ship,omitempty"`
		Year          string `json:"yea,omitemptyr"`
		Permalink     string `json:"permalink,omitempty"`
	} `json:"convict,omitempty"`
	Bankruptcy struct {
		BankruptDate string `json:"bankrupt_date,omitempty"`
		Year         string `json:"year,omitempty"`
		Permalink    string `json:"permalink,omitempty"`
	} `json:"bankruptcy,omitempty"`
	Marriage struct {
		MarriageDate  string `json:"marriage_date,omitempty"`
		SpouseName    string `json:"spouse_name,omitempty"`
		MarriagePlace string `json:"marriage_place,omitempty"`
		Year          string `json:"year,omitempty"`
		Permalink     string `json:"permalink,omitempty"`
	} `json:"marriage,omitempty"`
	Court struct {
		TrialDate    string `json:"trial_date,omitempty"`
		TrialOffence string `json:"trial_offence,omitempty"`
		Year         string `json:"year,omitempty"`
		Permalink    string `json:"permalink,omitempty"`
	} `json:"court,omitempty"`
	HealthWelfare struct {
		AdmissionDate string `json:"admission_date,omitempty"`
		Property      string `json:"property,omitempty"`
		Remarks       string `json:"remarks,omitempty"`
		Year          string `json:"year,omitempty"`
		Permalink     string `json:"permalink,omitempty"`
	} `json:"health-welfare,omitempty"`
	Census struct {
		CensusYear     string `json:"census_year,omitempty"`
		CensusPlace    string `json:"census_place,omitempty"`
		CensusChildren bool   `json:"census_children,omitempty"`
		Year           string `json:"year,omitempty"`
		Permalink      string `json:"permalink,omitempty"`
	} `json:"census,omitempty"`
}
