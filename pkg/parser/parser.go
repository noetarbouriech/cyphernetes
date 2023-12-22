package parser

import (
	"fmt"
	"log"
)

var Namespace string
var LogLevel string
var AllNamespaces bool

type Expression struct {
	Clauses []Clause
}

type Clause interface {
	isClause()
}

type MatchClause struct {
	Nodes         []*NodePattern
	Relationships []*Relationship
}

type SetClause struct {
	KeyValuePairs []*KeyValuePair
}

type KeyValuePair struct {
	Key   string
	Value interface{}
}

type Relationship struct {
	ResourceProperties *ResourceProperties
	Direction          Direction
	LeftNode           *NodePattern
	RightNode          *NodePattern
}

type NodeRelationshipList struct {
	Nodes         []*NodePattern
	Relationships []*Relationship
}

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	Both  Direction = "both"
	None  Direction = "none"
)

// type CreateClause struct {
// 	NodePattern         *NodePattern
// 	RelationshipPattern *RelationshipPattern
// }

// type SetClause struct {
// 	Identifier    string
// 	PropertyValue string
// }

type ReturnClause struct {
	JsonPaths []string
}

// type ReturnClause struct {
// 	ReturnItems []string
// }

type NodePattern struct {
	ResourceProperties *ResourceProperties
}

type ResourceProperties struct {
	Name       string
	Kind       string
	Properties *Properties
}

type Properties struct {
	PropertyList []*Property
}

type Property struct {
	Key string
	// Value is string int or bool
	Value interface{}
}

type JSONPathValueList struct {
	JSONPathValues []*JSONPathValue
}

type JSONPathValue struct {
	Value interface{}
}

// Implement isClause for all Clause types
func (m *MatchClause) isClause() {}

// func (c *CreateClause) isClause() {}
func (s *SetClause) isClause() {}

// func (d *DeleteClause) isClause() {}
func (r *ReturnClause) isClause() {}

var result *Expression

func ParseQuery(query string) (*Expression, error) {
	// Initialize the lexer with the query string
	lexer := NewLexer(query)

	// Call the parser function generated by goyacc
	// yyParse returns an int, with 0 indicating success
	if yyParse(lexer) != 0 {
		// Handle parsing error
		return nil, fmt.Errorf("parsing failed")
	}

	// Return the global result variable
	return result, nil
}

func logDebug(v ...interface{}) {
	if LogLevel == "debug" {
		log.Println(v...)
	}
}
