package schema

// WellAll schemas v0.1.0 (wellally.tech)

// ---------- Common ----------
type Coding struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display,omitempty"`
}

type CodeableConcept struct {
	Coding []Coding `json:"coding"`
	Text   string   `json:"text,omitempty"`
}

type Quantity struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ReferenceRange struct {
	Low  *Quantity `json:"low,omitempty"`
	High *Quantity `json:"high,omitempty"`
	Text string    `json:"text,omitempty"`
}

type Identifier struct {
	System      string           `json:"system"`
	Value       string           `json:"value"`
	Type        *CodeableConcept `json:"type,omitempty"`
	Period      *Period          `json:"period,omitempty"`
}

type HumanName struct {
	Family string   `json:"family"`
	Given  []string `json:"given"`
	Use    string   `json:"use,omitempty"`
	Prefix []string `json:"prefix,omitempty"`
	Suffix []string `json:"suffix,omitempty"`
}

type ContactPoint struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value"`
	Use    string `json:"use,omitempty"`
}

type Address struct {
	Line       []string `json:"line,omitempty"`
	City       string   `json:"city,omitempty"`
	State      string   `json:"state,omitempty"`
	PostalCode string   `json:"postalCode,omitempty"`
	Country    string   `json:"country,omitempty"`
}

type Period struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Modality struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display,omitempty"`
}

type Route struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display,omitempty"`
}

type SpecimenType struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display,omitempty"`
}

// ---------- Health Person ----------
type ClinicalSummary struct {
	Conditions           []CodeableConcept `json:"conditions,omitempty"`
	Allergies            []CodeableConcept `json:"allergies,omitempty"`
	BloodType            string            `json:"bloodType,omitempty"`
	PrimaryCareProvider  string            `json:"primaryCareProvider,omitempty"`
}

type HealthPerson struct {
	ID             string             `json:"id"`
	Name           []HumanName        `json:"name"`
	BirthDate      string             `json:"birthDate"`
	ResourceType   string             `json:"resourceType,omitempty"`
	Identifier     []Identifier       `json:"identifier,omitempty"`
	Gender         string             `json:"gender,omitempty"`
	Telecom        []ContactPoint     `json:"telecom,omitempty"`
	Address        []Address          `json:"address,omitempty"`
	MaritalStatus  *CodeableConcept   `json:"maritalStatus,omitempty"`
	Language       []string           `json:"language,omitempty"`
	ClinicalSummary *ClinicalSummary  `json:"clinicalSummary,omitempty"`
}

// ---------- Lab Report ----------
type LabResult struct {
	Code            CodeableConcept  `json:"code"`
	Value           any              `json:"value"` // Quantity | CodeableConcept | string
	ReferenceRange  *ReferenceRange  `json:"referenceRange,omitempty"`
	Interpretation  *string          `json:"interpretation,omitempty"`
	Method          *CodeableConcept `json:"method,omitempty"`
}

type Specimen struct {
	Type        *SpecimenType `json:"type,omitempty"`
	CollectedAt string        `json:"collectedAt,omitempty"`
}

type Facility struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type LabReport struct {
	ID        string       `json:"id"`
	PatientID string       `json:"patientId"`
	IssuedAt  string       `json:"issuedAt"`
	Results   []LabResult  `json:"results"`
	Facility  *Facility    `json:"facility,omitempty"`
	Panel     *CodeableConcept `json:"panel,omitempty"`
	Specimen  *Specimen    `json:"specimen,omitempty"`
}

// ---------- Imaging Report ----------
type Attachment struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type RadiationDose struct {
	CtdiVol_mGy *float64 `json:"ctdiVol_mGy,omitempty"`
	Dlp_mGy_cm  *float64 `json:"dlp_mGy_cm,omitempty"`
}

type Performer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type ImagingReport struct {
	ID              string      `json:"id"`
	PatientID       string      `json:"patientId"`
	Modality        Modality    `json:"modality"`
	BodySite        Coding      `json:"bodySite"`
	ReportedAt      string      `json:"reportedAt"`
	StudyInstanceUid string     `json:"studyInstanceUid,omitempty"`
	Performer       *Performer  `json:"performer,omitempty"`
	Findings        []string    `json:"findings,omitempty"`
	Impression      string      `json:"impression,omitempty"`
	RadiationDose   *RadiationDose `json:"radiationDose,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
}

// ---------- Medication ----------
type Dosage struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type MedicationRecord struct {
	ID          string           `json:"id"`
	PatientID   string           `json:"patientId"`
	Medication  Coding           `json:"medication"`
	Dosage      Dosage           `json:"dosage"`
	Route       Route            `json:"route"`
	StartDate   string           `json:"startDate"`
	Form        *Coding          `json:"form,omitempty"`
	Frequency   string           `json:"frequency,omitempty"`
	DurationDays *int            `json:"durationDays,omitempty"`
	EndDate     string           `json:"endDate,omitempty"`
	Indication  *CodeableConcept `json:"indication,omitempty"`
	Instructions string          `json:"instructions,omitempty"`
}

// ---------- Family Health ----------
type FamilyMember struct {
	ID               string            `json:"id"`
	RelationToProband string           `json:"relationToProband"`
	Sex              string            `json:"sex,omitempty"`
	BirthYear        *int              `json:"birthYear,omitempty"`
	Deceased         *bool             `json:"deceased,omitempty"`
	Conditions       []CodeableConcept `json:"conditions,omitempty"`
}

type FamilyHealthTree struct {
	ProbandID string         `json:"probandId"`
	Members   []FamilyMember `json:"members"`
}
