package couchdb

// Hector Oliveros - 2019
// hector.oliveros.leon@gmail.com

import (
	"base/config/ecodes"
	"base/pkg/cerror"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cabify/go-couchdb"
	"github.com/iancoleman/strcase"
	"net/http"
	"net/url"
	"reflect"
)

var DB *couchdb.DB

const couchDBDocTypeFieldName = "type"

type ClientConfig struct {
	Host               string
	User               string
	Pass               string
	PintOnCreate       bool
	DatabaseName       string
	CreateDBIfNotExist bool
}

func Post(doc interface{}) (id, rev string, err error) {
	ensuredDoc, err := ensureTypeField(doc)
	id, rev, err = DB.Post(ensuredDoc)
	return id, rev, err
}

type Row struct {
	ID, Key string
	Value   int
	Doc     interface{}
}

type ViewResult struct {
	TotalRows int `json:"total_rows"`
	Offset    int
	Rows      []Row
}

type ViewOpts struct {
	// Warning: ParseDoc Need optimize
	// Only use if you expect a few rows
	ParseDoc    bool     `json:"parseDoc" bson:"parseDoc"`
	Limit       int      `json:"limit" bson:"limit"`
	Offset      int      `json:"offset" bson:"offset"`
	IncludeDocs bool     `json:"includeDocs" bson:"includeDocs"`
	Key         string   `json:"key" bson:"key"`
	Keys        []string `json:"keys" bson:"keys"`
	StartKey    string   `json:"startKey" bson:"startKey"`
	EndKey      string   `json:"endKey" bson:"endKey"`
}

func viewOptsToOptions(vo ViewOpts) couchdb.Options {
	out := couchdb.Options{}
	if vo.Limit > 0 {
		out["limit"] = vo.Limit
	}
	if vo.Offset > 0 {
		out["offset"] = vo.Offset
	}
	if vo.IncludeDocs {
		out["include_docs"] = vo.IncludeDocs
	}
	if vo.Key != "" {
		out["key"] = vo.Key
	} else if vo.Keys != nil && len(vo.Keys) > 0 {
		out["keys"] = vo.Keys
	} else {
		if vo.StartKey != "" {
			out["startkey"] = vo.StartKey
		}
		if vo.EndKey != "" {
			out["endkey"] = vo.EndKey
		}
	}
	return out

}

type DDoc string
type ViewName string

func FindOneInView(ddoc DDoc, view ViewName, key string, docType interface{}) cerror.Error {
	result := ViewResult{}
	opts := ViewOpts{
		ParseDoc:    true,
		Limit:       1,
		IncludeDocs: true,
		Key:         key,
	}
	err := View(ddoc, view, &result, docType, opts)
	if err != nil {
		return cerror.From(err)
	}
	if len(result.Rows) == 0 {
		return cerror.From(ecodes.ErrNotFound.AddMeta("key", key).AddMeta(
			"ddoc", string(ddoc)).AddMeta("view", string(view)),
		)
	}
	docType = result.Rows[0].Doc
	return cerror.Nil
}

func View(ddoc DDoc, view ViewName, result *ViewResult, docType interface{}, vo ViewOpts) error {
	err := DB.View(string(ddoc), string(view), result, viewOptsToOptions(vo))
	if err != nil {
		return cerror.From(err)
	}

	if !vo.ParseDoc || docType == nil {
		return cerror.Nil
	}
	valueOf := reflect.ValueOf(docType)
	if valueOf.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("docType must be a pointer. got: %s", valueOf.Type().Kind().String())
	}

	for i := 0; i < len(result.Rows); i++ {
		inRec, err := json.Marshal(result.Rows[i].Doc)
		if err != nil {
			return cerror.From(err)
		}
		err = json.Unmarshal(inRec, docType)
		if err != nil {
			return cerror.From(err)
		}
		result.Rows[i].Doc = docType
	}
	return nil
}

func Configure(cfg ClientConfig) (*couchdb.DB, error) {
	err := validateCfg(cfg)
	if err != nil {
		return nil, err
	}
	c, err := createClient(cfg)
	if err != nil {
		return nil, err
	}
	if cfg.PintOnCreate {
		err := c.Ping() // trigger round trip
		if err != nil {
			return nil, fmt.Errorf("failed ping to db: %v", err)
		}
	}
	DB, err = c.EnsureDB(cfg.DatabaseName)
	if err != nil && cfg.CreateDBIfNotExist {
		DB, err = c.CreateDB(cfg.DatabaseName)
		if err != nil {
			return nil, fmt.Errorf("the database '%s' does not exist and can not be created", cfg.DatabaseName)
		}
	}
	return c.DB(cfg.DatabaseName), nil
}

func createClient(cfg ClientConfig) (*couchdb.Client, error) {
	httpClient := &http.Client{}
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("invalid url db: %s. %v", cfg.Host, err)
	}
	c := couchdb.NewClient(u, httpClient, couchdb.BasicAuth("admin", "123456"))
	return c, nil
}

func validateCfg(cfg ClientConfig) error {
	if cfg.Host == "" {
		return errors.New("empty host")
	}
	if cfg.User == "" && cfg.Pass != "" {
		return errors.New("empty user with not empty password")
	}
	if cfg.User != "" && cfg.Pass == "" {
		return errors.New("empty password with not empty user")
	}
	return nil
}

func getType(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)

	if valueOf.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(valueOf).Type().Name()
	} else {
		return valueOf.Type().Name()
	}
}

func ensureTypeField(doc interface{}) (map[string]interface{}, error) {
	in := &doc
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}
	_, haveTypeField := inInterface[couchDBDocTypeFieldName]
	if !haveTypeField {
		structName := strcase.ToLowerCamel(getType(doc))
		inInterface[couchDBDocTypeFieldName] = structName
	}
	return inInterface, nil
}
