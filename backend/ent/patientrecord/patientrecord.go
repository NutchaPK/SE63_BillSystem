// Code generated by entc, DO NOT EDIT.

package patientrecord

const (
	// Label holds the string label denoting the patientrecord type in the database.
	Label = "patientrecord"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIdcardnumber holds the string denoting the idcardnumber field in the database.
	FieldIdcardnumber = "idcardnumber"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldBloodtype holds the string denoting the bloodtype field in the database.
	FieldBloodtype = "bloodtype"
	// FieldDisease holds the string denoting the disease field in the database.
	FieldDisease = "disease"
	// FieldAllergic holds the string denoting the allergic field in the database.
	FieldAllergic = "allergic"
	// FieldPhonenumber holds the string denoting the phonenumber field in the database.
	FieldPhonenumber = "phonenumber"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldHome holds the string denoting the home field in the database.
	FieldHome = "home"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"

	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "gender"
	// EdgeMedicalrecordstaff holds the string denoting the medicalrecordstaff edge name in mutations.
	EdgeMedicalrecordstaff = "medicalrecordstaff"
	// EdgePrename holds the string denoting the prename edge name in mutations.
	EdgePrename = "prename"
	// EdgeHistorytaking holds the string denoting the historytaking edge name in mutations.
	EdgeHistorytaking = "historytaking"
	// EdgeTreatment holds the string denoting the treatment edge name in mutations.
	EdgeTreatment = "treatment"
	// EdgePatientrecordPatientrights holds the string denoting the patientrecordpatientrights edge name in mutations.
	EdgePatientrecordPatientrights = "PatientrecordPatientrights"

	// Table holds the table name of the patientrecord in the database.
	Table = "patientrecords"
	// GenderTable is the table the holds the gender relation/edge.
	GenderTable = "patientrecords"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the gender relation/edge.
	GenderColumn = "gender_id"
	// MedicalrecordstaffTable is the table the holds the medicalrecordstaff relation/edge.
	MedicalrecordstaffTable = "patientrecords"
	// MedicalrecordstaffInverseTable is the table name for the Medicalrecordstaff entity.
	// It exists in this package in order to avoid circular dependency with the "medicalrecordstaff" package.
	MedicalrecordstaffInverseTable = "medicalrecordstaffs"
	// MedicalrecordstaffColumn is the table column denoting the medicalrecordstaff relation/edge.
	MedicalrecordstaffColumn = "medicalrecordstaff_id"
	// PrenameTable is the table the holds the prename relation/edge.
	PrenameTable = "patientrecords"
	// PrenameInverseTable is the table name for the Prename entity.
	// It exists in this package in order to avoid circular dependency with the "prename" package.
	PrenameInverseTable = "prenames"
	// PrenameColumn is the table column denoting the prename relation/edge.
	PrenameColumn = "prefix"
	// HistorytakingTable is the table the holds the historytaking relation/edge.
	HistorytakingTable = "historytakings"
	// HistorytakingInverseTable is the table name for the Historytaking entity.
	// It exists in this package in order to avoid circular dependency with the "historytaking" package.
	HistorytakingInverseTable = "historytakings"
	// HistorytakingColumn is the table column denoting the historytaking relation/edge.
	HistorytakingColumn = "patientrecord_id"
	// TreatmentTable is the table the holds the treatment relation/edge.
	TreatmentTable = "treatments"
	// TreatmentInverseTable is the table name for the Treatment entity.
	// It exists in this package in order to avoid circular dependency with the "treatment" package.
	TreatmentInverseTable = "treatments"
	// TreatmentColumn is the table column denoting the treatment relation/edge.
	TreatmentColumn = "patientrecord_id"
	// PatientrecordPatientrightsTable is the table the holds the PatientrecordPatientrights relation/edge.
	PatientrecordPatientrightsTable = "patientrights"
	// PatientrecordPatientrightsInverseTable is the table name for the Patientrights entity.
	// It exists in this package in order to avoid circular dependency with the "patientrights" package.
	PatientrecordPatientrightsInverseTable = "patientrights"
	// PatientrecordPatientrightsColumn is the table column denoting the PatientrecordPatientrights relation/edge.
	PatientrecordPatientrightsColumn = "patientrecord_id"
)

// Columns holds all SQL columns for patientrecord fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIdcardnumber,
	FieldAge,
	FieldBirthday,
	FieldBloodtype,
	FieldDisease,
	FieldAllergic,
	FieldPhonenumber,
	FieldEmail,
	FieldHome,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Patientrecord type.
var ForeignKeys = []string{
	"gender_id",
	"medicalrecordstaff_id",
	"prefix",
}
