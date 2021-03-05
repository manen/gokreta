package endpoints

const (
	//! Base URLs

	// Global API, used to request the Schools
	Global = "https://kretaglobalmobileapi2.ekreta.hu:443/api/v2/"
	// IDP API, used to get a token
	IDP = "https://idp.e-kreta.hu/"
	// Admin API, used to get messages and stuff
	Admin = "https://eugyintezes.e-kreta.hu/api/v1/"
	// Base API, used to communicate with the
	// dedicated school server
	//
	// Important: Replace "$" with the institute ID
	Base = "https://$.ekreta.hu/"

	//! Sub-base URLs

	// Checker is a shorthand
	Checker = Base + "ellenorzo/V3/"

	// Own is a shorthand for the endpoints used
	// to request Evals, Absences, and Timetable
	Own = Checker + "Sajat/"

	// Communication is a shorthand
	//
	// what did you expect
	Communication = Admin + "/kommunikacio/"

	// MailboxItems is the shorthand for the
	// endpoints used to communicate with messages
	MailboxItems = Communication + "postaladaelemek/"

	// Employees is the shorthand for the endpoints
	// used to request data about the school's employees
	Employees = Admin + "/kreta/alkalmazottak/"

	//! Getter endpoints

	// OfficialSchools is the endpoint used to request the school list
	//
	// Use not recommended, requires a private API key, which
	// GoKreta does not include (DMCAs aren't fun)
	OfficialSchools = Global + "Institute"

	// Schools is an unofficial endpoint used by GoKreta
	// and the Filc mobile client. Does not require an API key.
	Schools = "https://raw.githubusercontent.com/filc/filc.github.io/master/docs/v2/school_list.json"

	// Token is the endpoint used to request a token
	Token = IDP + "connect/token"

	// OwnMessages is the endpoint used to request your own messages
	OwnMessages = MailboxItems + "sajat"
	// ReceivedMessages is the endpoint used to request the received messages
	ReceivedMessages = MailboxItems + "beerkezett"
	// SentMessages is the endpoint used to request the messages the user has written
	SentMessages = MailboxItems + "elkuldott"
	// DeletedMessages is the endpoint used to request the messages the user's deleted
	DeletedMessages = MailboxItems + "torolt"
	// UnreadMessageCount is the count of messages the user hasn't read
	UnreadMessageCount = MailboxItems + "olvasatlanokszama"
	// MessageContent requests the contents of a message
	//
	// Important: Replace "*" with the message ID
	MessageContent = MailboxItems + "*"

	// AddressedTypes is the endpoint used to request the types of people
	// you can send a message to
	//
	// (or something also no idea what this does)
	AddressedTypes = Admin + "adatszotarak/cimzetttipusok"

	// TeacherMessagingData is the endpoint used to request the teacher data?
	//
	// (guess what i also do not know what this does)
	TeacherMessagingData = Employees + "tanar"
	// HeadTeacherMessagingData is the endpoint used for...
	//
	// something i guess?
	HeadTeacherMessagingData = Employees + "osztalyfonokok"
	// BoardOfDirectorsMessagingData is... well like the other messaging data stuff but board of director i guess idk
	BoardOfDirectorsMessagingData = Employees + "igazgatosag"
	// AdminMessagingData is messaging data for the administrator idk
	AdminMessagingData = Employees + "adminisztrator"
	// MessagableGroups is the groups the user can message
	MessagableGroups = Communication + "tanoraicsoportok/cimezheto"
	// MessagableSzmkRepresentatives is the Szmk representatives you can message
	MessagableSzmkRepresentatives = Communication + "szmkkepviselok/cimezheto"
	// DownloadFile is the endpoint used to download a file
	//
	// Important: Replace "*" with the file ID
	DownloadFile = Admin + "dokumentumok/uzenetek/*"

	// HeadTeacherData is the endpoint used to request the head teacher's data
	//
	// Important: Replace "*" with the teacher ID
	HeadTeacherData = Checker + "Felhasznalok/Alkalmazottak/Tanarok/Osztalyfonokok?Uids=*"

	// Grades is the endpoint used to request evaluations
	Grades = Own + "Ertekelesek"
	// Absences is the endpoint used to request absences
	Absences = Own + "Mulasztasok"
	// Timetable is the endpoint used to request the user's timetable
	Timetable = Own + "OrarendElemek"
	// Notes is the endpoint used to request notes
	Notes = Own + "Feljegyzesek"
	// Notices is the endpoint used to request notices
	Notices = Own + "FaliujsagElemek"
	// Tests is the endpoint used to request announced Tests
	Tests = Own + "BejelentettSzamonkeresek"
	// Groups is the endpoint used to request class Groups
	Groups = Own + "OsztalyCsoportok"
	// UserData is the endpoint used to request the user's data
	UserData = Own + "TanuloAdatlap"
	// YearOrder is the endpoint used to request the current year's order
	//
	// (I genuinely don't know what this one does lmao)
	YearOrder = Own + "TanevRendjeElemek"
	// Homework is the endpoint used to request the homework
	Homework = Own + "HaziFeladatok"
	// HomeworkComments is the endpoint used to request comments on an individual homework
	//
	// Important: Replace "*" with the homework ID
	HomeworkComments = Homework + "*/Kommentek"

	//! Setter/etc endpoints

	// HomeworkComment is the POST endpoint used to comment on homework
	//
	// The body should look something like: (JavaScript snippet)
	//  { HaziFeladatUid: homeworkId, FeladatSzovege: commentText }
	HomeworkComment = Own + "Orak/TanitasiOrak/HaziFeladatok/Kommentek"
)
