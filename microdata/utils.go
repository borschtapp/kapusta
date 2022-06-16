package microdata

import "golang.org/x/net/html"

// walkNodes traverses the node tree executing the given functions.
func walkNodes(n *html.Node, f func(*html.Node)) {
	if n != nil {
		f(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walkNodes(c, f)
		}
	}
}

// getAttr returns the value associated with the given attribute from the given node.
func getAttr(attribute string, node *html.Node) (string, bool) {
	for _, attr := range node.Attr {
		if attribute == attr.Key {
			return attr.Val, true
		}
	}
	return "", false
}

// checkAttr returns true if the given node has the given attribute.
func checkAttr(attribute, expValue string, node *html.Node) bool {
	for _, a := range node.Attr {
		if a.Key == attribute && a.Val == expValue {
			return true
		}
	}
	return false
}
