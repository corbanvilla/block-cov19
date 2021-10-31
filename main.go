package main

// Block
// Base block chain requirement
type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string
}

type Vaccine struct {
	Doctor		string
	Country		string
	CountrySig	string

	PatientId	int
	PatientSig	string

	VaccineType	string
	Location	string

	BatchId		int
	VaccineId	int

	Block
}

type PCRSample struct {
	Technician	string

	Country		string
	CountrySig	string

	PatientId	int
	PatientSig	string

	Location	string

	TestId		int
	BatchId		int

	Block
}

type PCRResult struct {
	Country		string
	CountrySig	string

	PatientId	int

	Location	string

	TestId		int
	BatchId		int

	TestResult	string

	Block
}

type FaceArray struct {
	Technician	string

	Country		string
	CountrySig	string

	PatientId	int
	PatientSig	string

	FaceEncoding	string

	Block
}

type CitizenIdentity struct {
	Country		string
	CountrySig	string

	PatientId	int
	PatientSig	string
	PatientKey	string

	Block
}

type CountryIdentity struct {
	Country		string
	CountrySig	string
	CountryKey	string

	Block
}
