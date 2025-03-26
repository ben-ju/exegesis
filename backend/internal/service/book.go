package service

import (
	"net/http"

	"github.com/ben-ju/exegesis/internal/repository"
	"github.com/ben-ju/exegesis/internal/router"
)

// StartBooks contient les livres possibles pour le début.
var StartBooks = []string{"Genesis", "Genèse"}

// EndBooks contient les livres possibles pour la fin.
var EndBooks = []string{"Revelation", "Apocalypse"}

// BooksEN liste les livres en anglais.
var BooksEN = []string{
	// Pentateuque
	"Genesis", "Exodus", "Leviticus", "Numbers", "Deuteronomy",
	// Historiques
	"Joshua", "Judges", "Ruth", "1 Samuel", "2 Samuel", "1 Kings", "2 Kings",
	"1 Chronicles", "2 Chronicles", "Ezra", "Nehemiah", "Tobit", "Judith",
	"Esther", "1 Maccabees", "2 Maccabees",
	// Poétiques
	"Job", "Psalms", "Proverbs", "Ecclesiastes", "Song of Songs", "Wisdom of Solomon",
	"Sirach",
	// Prophétiques
	"Isaiah", "Jeremiah", "Lamentations", "Baruch", "Ezekiel", "Daniel", "Hosea",
	"Joel", "Amos", "Obadiah", "Jonah", "Micah", "Nahum", "Habakkuk", "Zephaniah",
	"Haggai", "Zechariah", "Malachi",
	// NT
	"Matthew", "Mark", "Luke", "John", "Acts", "Romans", "1 Corinthians", "2 Corinthians",
	"Galatians", "Ephesians", "Philippians", "Colossians", "1 Thessalonians",
	"2 Thessalonians", "1 Timothy", "2 Timothy", "Titus", "Philemon", "Hebrews",
	"James", "1 Peter", "2 Peter", "1 John", "2 John", "3 John", "Jude", "Revelation",
}

// BooksFR liste les livres en français.
var BooksFR = []string{
	// Pentateuque
	"Genèse", "Exode", "Lévitique", "Nombres", "Deutéronome",
	// Historiques
	"Josué", "Juges", "Ruth", "1 Samuel", "2 Samuel", "1 Rois", "2 Rois",
	"1 Chroniques", "2 Chroniques", "Esdras", "Néhémie", "Tobie", "Judith",
	"Esther", "1 Maccabées", "2 Maccabées",
	// Poétiques
	"Job", "Psaumes", "Psaume", "Proverbes", "Ecclésiaste", "Cantique des Cantiques",
	"Cantique", "Sagesse", "Siracide",
	// Prophétiques
	"Isaïe", "Esaïe", "Jérémie", "Lamentations", "Baruch", "Ézéchiel", "Ezéchiel",
	"Daniel", "Osée", "Joël", "Amos", "Abdias", "Jonas", "Michée", "Nahum",
	"Habacuc", "Sophonie", "Aggée", "Zacharie", "Malachie",
	// NT
	"Matthieu", "Marc", "Luc", "Jean", "Actes", "Romains", "1 Corinthiens",
	"2 Corinthiens", "Galates", "Ephésiens", "Éphésiens", "Philippiens",
	"Colossiens", "1 Thessaloniciens", "2 Thessaloniciens", "1 Timothée", "2 Timothée",
	"Tite", "Philémon", "Hébreux", "Jacques", "1 Pierre", "2 Pierre", "1 Jean",
	"2 Jean", "3 Jean", "Jude", "Apocalypse",
}

// FrenchToEnglish mappe les noms de livres en français vers leur équivalent anglais.
var FrenchToEnglish = map[string]string{
	"Genèse": "Genesis", "Genese": "Genesis", "Génèse": "Genesis",
	"Exode":     "Exodus",
	"Lévitique": "Leviticus", "Levitique": "Leviticus",
	"Nombres":      "Numbers",
	"Deutéronome":  "Deuteronomy",
	"Josué":        "Joshua",
	"Juges":        "Judges",
	"Ruth":         "Ruth",
	"1 Samuel":     "1 Samuel",
	"2 Samuel":     "2 Samuel",
	"1 Rois":       "1 Kings",
	"2 Rois":       "2 Kings",
	"1 Chroniques": "1 Chronicles",
	"2 Chroniques": "2 Chronicles",
	"Esdras":       "Ezra",
	"Néhémie":      "Nehemiah", "Nehemie": "Nehemiah",
	"Tobie":       "Tobit",
	"Judith":      "Judith",
	"Esther":      "Esther",
	"1 Maccabées": "1 Maccabees", "1 Maccabees": "1 Maccabees",
	"2 Maccabées": "2 Maccabees", "2 Maccabees": "2 Maccabees",
	"Job":     "Job",
	"Psaumes": "Psalms", "Psaume": "Psalms",
	"Proverbes":   "Proverbs",
	"Ecclésiaste": "Ecclesiastes", "Ecclesiaste": "Ecclesiastes",
	"Cantique des Cantiques": "Song of Solomon", "Cantique": "Song of Solomon",
	"Sagesse":  "Wisdom of Solomon",
	"Siracide": "Sirach",
	"Isaïe":    "Isaiah", "Esaïe": "Isaiah", "Ésaïe": "Isaiah",
	"Jérémie":      "Jeremiah",
	"Lamentations": "Lamentations",
	"Baruch":       "Baruch",
	"Ézéchiel":     "Ezekiel", "Ezéchiel": "Ezekiel", "Ezechiel": "Ezekiel",
	"Daniel": "Daniel",
	"Osée":   "Hosea",
	"Joël":   "Joel", "Joel": "Joel",
	"Amos":          "Amos",
	"Abdias":        "Obadiah",
	"Jonas":         "Jonah",
	"Michée":        "Micah",
	"Nahum":         "Nahum",
	"Habacuc":       "Habakkuk",
	"Sophonie":      "Zephaniah",
	"Aggée":         "Haggai",
	"Zacharie":      "Zechariah",
	"Malachie":      "Malachi",
	"Matthieu":      "Matthew",
	"Marc":          "Mark",
	"Luc":           "Luke",
	"Jean":          "John",
	"Actes":         "Acts",
	"Romains":       "Romans",
	"1 Corinthiens": "1 Corinthians",
	"2 Corinthiens": "2 Corinthians",
	"Galates":       "Galatians",
	"Éphésiens":     "Ephesians", "Ephésiens": "Ephesians", "Ephesiens": "Ephesians", "Éphesiens": "Ephesians",
	"Philippiens":       "Philippians",
	"Colossiens":        "Colossians",
	"1 Thessaloniciens": "1 Thessalonians",
	"2 Thessaloniciens": "2 Thessalonians",
	"1 Timothée":        "1 Timothy", "1 Timothee": "1 Timothy",
	"2 Timothee": "2 Timothy", "2 Timothée": "2 Timothy",
	"Tite":     "Titus",
	"Philémon": "Philemon", "Philemon": "Philemon",
	"Hébreux": "Hebrews", "Hebreux": "Hebrews",
	"Jacques":    "James",
	"1 Pierre":   "1 Peter",
	"2 Pierre":   "2 Peter",
	"1 Jean":     "1 John",
	"2 Jean":     "2 John",
	"3 Jean":     "3 John",
	"Jude":       "Jude",
	"Apocalypse": "Revelation",
}

// Book représente un livre de la Bible.
type Book struct {
	Title              string `json:"title"`
	Abbreviation       string `json:"abbreviation"`
	IsDeuterocanonical int    `json:"is_deuterocanonical"`
	IsOldTestament     int    `json:"is_old_testament"`
	IsNewTestament     int    `json:"is_new_testament"`
}

// SeedBooks est la liste de livres à insérer initialement.
var SeedBooks = []Book{
	// Ancien Testament
	{Title: "Genesis", Abbreviation: "Gen", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Exodus", Abbreviation: "Exod", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Leviticus", Abbreviation: "Lev", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Numbers", Abbreviation: "Num", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Deuteronomy", Abbreviation: "Deut", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Joshua", Abbreviation: "Josh", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Judges", Abbreviation: "Judg", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ruth", Abbreviation: "Ruth", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Samuel", Abbreviation: "1 Sam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Samuel", Abbreviation: "2 Sam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Kings", Abbreviation: "1 Kgs", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Kings", Abbreviation: "2 Kgs", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Chronicles", Abbreviation: "1 Chr", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Chronicles", Abbreviation: "2 Chr", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ezra", Abbreviation: "Ezra", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Nehemiah", Abbreviation: "Neh", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Tobit", Abbreviation: "Tobit", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Judith", Abbreviation: "Judith", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Esther", Abbreviation: "Est", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Maccabees", Abbreviation: "1 Macc", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Maccabees", Abbreviation: "2 Macc", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Job", Abbreviation: "Job", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Psalms", Abbreviation: "Ps", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Proverbs", Abbreviation: "Prov", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ecclesiastes", Abbreviation: "Eccl", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Songs of Songs", Abbreviation: "Song", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Wisdom of Solomon", Abbreviation: "Wisd", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Sirach", Abbreviation: "Sir", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Isaiah", Abbreviation: "Isa", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Jeremiah", Abbreviation: "Jer", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Lamentations", Abbreviation: "Lam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Baruch", Abbreviation: "Baru", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ezekiel", Abbreviation: "Ezek", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Daniel", Abbreviation: "Dan", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Hosea", Abbreviation: "Hos", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Joel", Abbreviation: "Joel", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Amos", Abbreviation: "Amos", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Obadiah", Abbreviation: "Obad", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Jonah", Abbreviation: "Jonah", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Micah", Abbreviation: "Mic", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Nahum", Abbreviation: "Nah", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Habakkuk", Abbreviation: "Hab", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Zephaniah", Abbreviation: "Zeph", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Haggai", Abbreviation: "Hag", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Zechariah", Abbreviation: "Zech", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Malachi", Abbreviation: "Mal", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},

	// Nouveau Testament
	{Title: "Matthew", Abbreviation: "Matt", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Mark", Abbreviation: "Mark", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Luke", Abbreviation: "Luke", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "John", Abbreviation: "John", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Acts", Abbreviation: "Acts", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Romans", Abbreviation: "Rom", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Corinthians", Abbreviation: "1 Cor", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Corinthians", Abbreviation: "2 Cor", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Galatians", Abbreviation: "Gal", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Ephesians", Abbreviation: "Eph", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Philippians", Abbreviation: "Phil", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Colossians", Abbreviation: "Col", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Thessalonians", Abbreviation: "1 Thess", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Thessalonians", Abbreviation: "2 Thess", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Timothy", Abbreviation: "1 Tim", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Timothy", Abbreviation: "2 Tim", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Titus", Abbreviation: "Titus", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Philemon", Abbreviation: "Philem", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Hebrews", Abbreviation: "Heb", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "James", Abbreviation: "Jas", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Peter", Abbreviation: "1 Pet", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Peter", Abbreviation: "2 Pet", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 John", Abbreviation: "1 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 John", Abbreviation: "2 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "3 John", Abbreviation: "3 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Jude", Abbreviation: "Jude", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Revelation", Abbreviation: "Rev", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
}

type BookRoutable struct {
	repo *repository.BookRepository
}

func NewBookRoutable(repo *repository.BookRepository) *BookRoutable {
	return &BookRoutable{repo: repo}
}

func (br *BookRoutable) RegisterRoutes(r *router.Router) {
	r.HandleFunc("/versions", br.listVersions)
}

func (br *BookRoutable) listVersions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TOTO"))
	// versions, err := br.repo.ListVersions()
	// if err != nil {
	// 	http.Error(w, "Error fetching versions", http.StatusInternalServerError)
	// 	return
	// }
	// json.NewEncoder(w).Encode(versions)
}
