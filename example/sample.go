package example

import (
	"regexp"
	"strings"
)

// VuFindBibliographicIndex generated from 405560333c6c79e90b6b07ff5c916c353741e92d.
type VuFindBibliographicIndex struct {
	Version              string   `json:"_version_"`
	ID                   string   `json:"id"`
	Fullrecord           string   `json:"fullrecord"`
	Itemdata             string   `json:"itemdata"`
	MARCError            []string `json:"marc_error"`
	Allfields            []string `json:"allfields"`
	AllfieldsUnstemmed   []string `json:"allfields_unstemmed"`
	Fulltext             string   `json:"fulltext"`
	FulltextUnstemmed    string   `json:"fulltext_unstemmed"`
	Spelling             []string `json:"spelling"`
	SpellingShingle      []string `json:"spellingShingle"`
	AccessFacet          string   `json:"access_facet"`
	MegaCollection       []string `json:"mega_collection"`
	RecordID             string   `json:"record_id"`
	SourceID             string   `json:"source_id"`
	Institution          []string `json:"institution"`
	Collection           []string `json:"collection"`
	Building             []string `json:"building"`
	Language             []string `json:"language"`
	Format               []string `json:"format"`
	Signatur             []string `json:"signatur"`
	Barcode              []string `json:"barcode"`
	RSN                  []string `json:"rsn"`
	Purchase             []string `json:"purchase"`
	Timecode             []string `json:"timecode"`
	MiscDech1            []string `json:"misc_dech1"`
	FincClassFacet       []string `json:"finc_class_facet"`
	ZDB                  string   `json:"zdb"`
	AuthorID             []string `json:"author_id"`
	AuthorRef            []string `json:"author_ref"`
	AuthorCorporateRef   []string `json:"author_corporate_ref"`
	AuthorStr            string   `json:"authorStr"`
	AuthorFacet          []string `json:"author_facet"`
	AuthorSort           string   `json:"author_sort"`
	Author               []string `json:"author"`
	AuthorOrig           []string `json:"author_orig"`
	AuthorRole           []string `json:"author_role"`
	Author2              []string `json:"author2"`
	Author2Orig          []string `json:"author2_orig"`
	Author2Role          []string `json:"author2_role"`
	AuthorCorporate      []string `json:"author_corporate"`
	AuthorCorporateOrig  []string `json:"author_corporate_orig"`
	AuthorCorporateRole  []string `json:"author_corporate_role"`
	AuthorCorporate2     []string `json:"author_corporate2"`
	AuthorCorporate2Orig []string `json:"author_corporate2_orig"`
	AuthorCorporate2Role []string `json:"author_corporate2_role"`
	AuthorAdditional     []string `json:"author_additional"`
	Title                string   `json:"title"`
	TitlePart            string   `json:"title_part"`
	TitleSub             string   `json:"title_sub"`
	TitleShort           string   `json:"title_short"`
	TitleFull            string   `json:"title_full"`
	TitleFullUnstemmed   string   `json:"title_full_unstemmed"`
	TitleFullStr         string   `json:"title_fullStr"`
	TitleAuth            string   `json:"title_auth"`
	TitleAlt             []string `json:"title_alt"`
	TitleOld             []string `json:"title_old"`
	TitleNew             []string `json:"title_new"`
	TitleUniform         string   `json:"title_uniform"`
	TitleSort            string   `json:"title_sort"`
	TitleOrig            string   `json:"title_orig"`
	Physical             []string `json:"physical"`
	Publisher            []string `json:"publisher"`
	PublisherStr         []string `json:"publisherStr"`
	PublishDate          []string `json:"publishDate"`
	PublishDateSort      string   `json:"publishDateSort"`
	PublishPlace         []string `json:"publishPlace"`
	Imprint              string   `json:"imprint"`
	Edition              string   `json:"edition"`
	Description          string   `json:"description"`
	Contents             []string `json:"contents"`
	URL                  []string `json:"url"`
	Thumbnail            string   `json:"thumbnail"`
	LCCN                 string   `json:"lccn"`
	Ctrlnum              []string `json:"ctrlnum"`
	URN                  []string `json:"urn"`
	ISBN                 []string `json:"isbn"`
	ISSN                 []string `json:"issn"`
	ISMN                 []string `json:"ismn"`
	OCLCNum              []string `json:"oclc_num"`
	Callnumberfirst      string   `json:"callnumber-first"`
	Callnumbersubject    string   `json:"callnumber-subject"`
	Callnumberlabel      string   `json:"callnumber-label"`
	Callnumbersort       string   `json:"callnumber-sort"`
	Callnumberraw        []string `json:"callnumber-raw"`
	Callnumbersearch     []string `json:"callnumber-search"`
	Deweyhundreds        []string `json:"dewey-hundreds"`
	Deweytens            []string `json:"dewey-tens"`
	Deweyones            []string `json:"dewey-ones"`
	Deweyfull            []string `json:"dewey-full"`
	Deweysort            string   `json:"dewey-sort"`
	Deweyraw             []string `json:"dewey-raw"`
	Deweysearch          []string `json:"dewey-search"`
	DateSpan             []string `json:"dateSpan"`
	Series               []string `json:"series"`
	Series2              []string `json:"series2"`
	SeriesOrig           []string `json:"series_orig"`
	Topic                []string `json:"topic"`
	TopicID              []string `json:"topic_id"`
	TopicRef             []string `json:"topic_ref"`
	TopicUnstemmed       []string `json:"topic_unstemmed"`
	TopicFacet           []string `json:"topic_facet"`
	TopicBrowse          []string `json:"topic_browse"`
	AuthorBrowse         []string `json:"author_browse"`
	Genre                []string `json:"genre"`
	GenreFacet           []string `json:"genre_facet"`
	Geographic           []string `json:"geographic"`
	GeographicFacet      []string `json:"geographic_facet"`
	GeogrCode            []string `json:"geogr_code"`
	GeogrCodePerson      []string `json:"geogr_code_person"`
	Era                  []string `json:"era"`
	EraFacet             []string `json:"era_facet"`
	Footnote             []string `json:"footnote"`
	DissertationNote     []string `json:"dissertation_note"`
	PerformerNote        []string `json:"performer_note"`
	Illustrated          string   `json:"illustrated"`
	LongLat              string   `json:"long_lat"`
	MusicHeading         []string `json:"music_heading"`
	MusicHeadingBrowse   []string `json:"music_heading_browse"`
	FilmHeading          []string `json:"film_heading"`
	RVKFacet             []string `json:"rvk_facet"`
	RVKLabel             []string `json:"rvk_label"`
	RVKPath              []string `json:"rvk_path"`
	ContainerTitle       string   `json:"container_title"`
	ContainerVolume      string   `json:"container_volume"`
	ContainerIssue       string   `json:"container_issue"`
	ContainerStartPage   string   `json:"container_start_page"`
	ContainerReference   string   `json:"container_reference"`
	MultipartSet         string   `json:"multipart_set"`
	MultipartLink        []string `json:"multipart_link"`
	MultipartPart        []string `json:"multipart_part"`
	Hierarchytype        string   `json:"hierarchytype"`
	HierarchyTopID       []string `json:"hierarchy_top_id"`
	HierarchyTopTitle    []string `json:"hierarchy_top_title"`
	HierarchyParentID    []string `json:"hierarchy_parent_id"`
	HierarchyParentTitle []string `json:"hierarchy_parent_title"`
	HierarchySequence    []string `json:"hierarchy_sequence"`
	IsHierarchyID        string   `json:"is_hierarchy_id"`
	IsHierarchyTitle     string   `json:"is_hierarchy_title"`
	TitleInHierarchy     []string `json:"title_in_hierarchy"`
	HierarchyBrowse      []string `json:"hierarchy_browse"`
	Recordtype           string   `json:"recordtype"`
	FirstIndexed         string   `json:"first_indexed"`
	LastIndexed          string   `json:"last_indexed"`

	dynamicFields []struct {
		Key    string
		Values []string
	}
}

// allowedDynamicFieldName returns true, if the name of the field matches
// one of the dynamic field patterns.
func (v VuFindBibliographicIndex) allowedDynamicFieldName(k string) (ok bool, err error) {
	return WildcardMatch(k, []string{
		"callnumber_*",
		"barcode_*",
		"misc_*",
		"branch_*",
		"collcode_*",
		"format_*",
		"facet_*",
		"local_heading_facet_*",
		"local_heading_*",
		"local_class_*",
		"udk_raw_*",
		"udk_facet_*",
		"date_*",
		"*_date",
		"*_date_mv",
		"*_isn",
		"*_isn_mv",
		"*_str",
		"*_str_mv",
		"*_txt",
		"*_txt_mv",
		"*_txtF",
		"*_txtF_mv",
		"*_txtP",
		"*_txtP_mv",
		"*_random",
		"*_boolean",
	})
}

// WildcardMatch returns true, if the wildcard covers a given string s. If the
// wildcard is invalid, an error is returned.
func WildcardMatch(s, wildcards []string) (bool, error) {
	for _, w := range wildcards {
		p := strings.Replace(w, "*", ".*", -1)
		p = "^" + p + "$"
		r, err := regexp.Compile(p)
		if err != nil {
			return false, err
		}
		if r.MatchString(s) {
			return true, nil
		}
	}
	return false, nil
}
