// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbolsearch

import "fmt"

var Content = fmt.Sprintf(`// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated with go generate -run gen_query.go. DO NOT EDIT.

package symbolsearch

// querySearchSymbol is used when the search query is only one word, with no dots.
// In this case, the word must match a symbol name and ranking is completely
// determined by the path_tokens.
%s

// querySearchPackageDotSymbol is used when the search query is one element
// containing a dot, where the first part is assumed to be the package name and
// the second the symbol name. For example, "sql.DB" or "sql.DB.Begin".
%s

// querySearchMultiWordExact is used when the search query is multiple elements.
%s

// querySearchMultiWordOr is used when the search query is multiple elements.
%s

// queryMatchingSymbolIDsSymbol is used to find the matching symbol
// ids when the SearchType is SearchTypeSymbol.
%s

// queryMatchingSymbolIDsPackageDotSymbol is used to find the matching symbol
// ids when the SearchType is SearchTypePackageDotSymbol.
%s

// queryMatchingSymbolIDsMultiWordExact is used to find the matching symbol ids when
// the SearchType is SearchTypeMultiWordExact.
%s

// queryMatchingSymbolIDsMultiWordOr is used to find the matching symbol ids when
// the SearchType is SearchTypeMultiWordOr.
%s
`,
	formatQuery("querySearchSymbol", Query(SearchTypeSymbol)),
	formatQuery("querySearchPackageDotSymbol", Query(SearchTypePackageDotSymbol)),
	formatQuery("querySearchMultiWordOr", Query(SearchTypeMultiWordOr)),
	formatQuery("querySearchMultiWordExact", Query(SearchTypeMultiWordExact)),
	formatQuery("queryMatchingSymbolIDsSymbol", MatchingSymbolIDsQuery(SearchTypeSymbol)),
	formatQuery("queryMatchingSymbolIDsPackageDotSymbol", MatchingSymbolIDsQuery(SearchTypePackageDotSymbol)),
	formatQuery("queryMatchingSymbolIDsMultiWordOr", MatchingSymbolIDsQuery(SearchTypeMultiWordOr)),
	formatQuery("queryMatchingSymbolIDsMultiWordExact", MatchingSymbolIDsQuery(SearchTypeMultiWordExact)))

func formatQuery(name, query string) string {
	return fmt.Sprintf("const %s = `%s`", name, query)
}
