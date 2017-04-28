package expression

import . "github.com/pingcap/check"

func testPathExprBodyRe(c *C) {
	var pathExpr = "$.key1[3][6].key2"
	matches := pathExprBodyRe.FindStringSubmatchIndex(pathExpr)
	c.Assert(matches[0], Equals, 0)
}
