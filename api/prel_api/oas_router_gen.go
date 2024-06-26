// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleGetRequest([0]string{}, elemIsEscaped, w, r)
				default:
					s.notAllowed(w, r, "GET")
				}

				return
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				origElem := elem
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "dmin/"
					origElem := elem
					if l := len("dmin/"); len(elem) >= l && elem[0:l] == "dmin/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "iam-role-filtering"
						origElem := elem
						if l := len("iam-role-filtering"); len(elem) >= l && elem[0:l] == "iam-role-filtering" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAdminIamRoleFilteringGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					case 'r': // Prefix: "request"
						origElem := elem
						if l := len("request"); len(elem) >= l && elem[0:l] == "request" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAdminRequestGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					case 's': // Prefix: "setting"
						origElem := elem
						if l := len("setting"); len(elem) >= l && elem[0:l] == "setting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAdminSettingGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					case 'u': // Prefix: "user"
						origElem := elem
						if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAdminUserGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "pi/"
					origElem := elem
					if l := len("pi/"); len(elem) >= l && elem[0:l] == "pi/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "i"
						origElem := elem
						if l := len("i"); len(elem) >= l && elem[0:l] == "i" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "am-role"
							origElem := elem
							if l := len("am-role"); len(elem) >= l && elem[0:l] == "am-role" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '-': // Prefix: "-filtering-rules"
								origElem := elem
								if l := len("-filtering-rules"); len(elem) >= l && elem[0:l] == "-filtering-rules" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "GET":
										s.handleAPIIamRoleFilteringRulesGetRequest([0]string{}, elemIsEscaped, w, r)
									case "POST":
										s.handleAPIIamRoleFilteringRulesPostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "GET,POST")
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"
									origElem := elem
									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "ruleID"
									// Leaf parameter
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleAPIIamRoleFilteringRulesRuleIDDeleteRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, "DELETE")
										}

										return
									}

									elem = origElem
								}

								elem = origElem
							case 's': // Prefix: "s"
								origElem := elem
								if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleAPIIamRolesGetRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}

								elem = origElem
							}

							elem = origElem
						case 'n': // Prefix: "nvitations"
							origElem := elem
							if l := len("nvitations"); len(elem) >= l && elem[0:l] == "nvitations" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAPIInvitationsPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					case 'r': // Prefix: "requests"
						origElem := elem
						if l := len("requests"); len(elem) >= l && elem[0:l] == "requests" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleAPIRequestsGetRequest([0]string{}, elemIsEscaped, w, r)
							case "POST":
								s.handleAPIRequestsPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET,POST")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							origElem := elem
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "requestID"
							// Leaf parameter
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "DELETE":
									s.handleAPIRequestsRequestIDDeleteRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								case "PATCH":
									s.handleAPIRequestsRequestIDPatchRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "DELETE,PATCH")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					case 's': // Prefix: "settings"
						origElem := elem
						if l := len("settings"); len(elem) >= l && elem[0:l] == "settings" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PATCH":
								s.handleAPISettingsPatchRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PATCH")
							}

							return
						}

						elem = origElem
					case 'u': // Prefix: "users"
						origElem := elem
						if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleAPIUsersGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							origElem := elem
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "userID"
							// Leaf parameter
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PATCH":
									s.handleAPIUsersUserIDPatchRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PATCH")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'u': // Prefix: "uth/google/callback"
					origElem := elem
					if l := len("uth/google/callback"); len(elem) >= l && elem[0:l] == "uth/google/callback" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAuthGoogleCallbackGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "health"
				origElem := elem
				if l := len("health"); len(elem) >= l && elem[0:l] == "health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHealthGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'r': // Prefix: "request"
				origElem := elem
				if l := len("request"); len(elem) >= l && elem[0:l] == "request" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleRequestGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '-': // Prefix: "-form"
					origElem := elem
					if l := len("-form"); len(elem) >= l && elem[0:l] == "-form" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRequestFormGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "requestID"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleRequestRequestIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "sign"
				origElem := elem
				if l := len("sign"); len(elem) >= l && elem[0:l] == "sign" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "in"
					origElem := elem
					if l := len("in"); len(elem) >= l && elem[0:l] == "in" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleSigninPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'o': // Prefix: "out"
					origElem := elem
					if l := len("out"); len(elem) >= l && elem[0:l] == "out" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleSignoutPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = "Get"
					r.summary = "display top page"
					r.operationID = ""
					r.pathPattern = "/"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				origElem := elem
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "dmin/"
					origElem := elem
					if l := len("dmin/"); len(elem) >= l && elem[0:l] == "dmin/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "iam-role-filtering"
						origElem := elem
						if l := len("iam-role-filtering"); len(elem) >= l && elem[0:l] == "iam-role-filtering" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "AdminIamRoleFilteringGet"
								r.summary = "return iam role filtering page"
								r.operationID = ""
								r.pathPattern = "/admin/iam-role-filtering"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'r': // Prefix: "request"
						origElem := elem
						if l := len("request"); len(elem) >= l && elem[0:l] == "request" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "AdminRequestGet"
								r.summary = "return admin request page"
								r.operationID = ""
								r.pathPattern = "/admin/request"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 's': // Prefix: "setting"
						origElem := elem
						if l := len("setting"); len(elem) >= l && elem[0:l] == "setting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "AdminSettingGet"
								r.summary = "return admin setting page"
								r.operationID = ""
								r.pathPattern = "/admin/setting"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'u': // Prefix: "user"
						origElem := elem
						if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "AdminUserGet"
								r.summary = "return admin user page"
								r.operationID = ""
								r.pathPattern = "/admin/user"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "pi/"
					origElem := elem
					if l := len("pi/"); len(elem) >= l && elem[0:l] == "pi/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'i': // Prefix: "i"
						origElem := elem
						if l := len("i"); len(elem) >= l && elem[0:l] == "i" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "am-role"
							origElem := elem
							if l := len("am-role"); len(elem) >= l && elem[0:l] == "am-role" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '-': // Prefix: "-filtering-rules"
								origElem := elem
								if l := len("-filtering-rules"); len(elem) >= l && elem[0:l] == "-filtering-rules" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										r.name = "APIIamRoleFilteringRulesGet"
										r.summary = "return iam role filtering rules"
										r.operationID = ""
										r.pathPattern = "/api/iam-role-filtering-rules"
										r.args = args
										r.count = 0
										return r, true
									case "POST":
										r.name = "APIIamRoleFilteringRulesPost"
										r.summary = "post iam role filtering rule"
										r.operationID = ""
										r.pathPattern = "/api/iam-role-filtering-rules"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"
									origElem := elem
									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "ruleID"
									// Leaf parameter
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = "APIIamRoleFilteringRulesRuleIDDelete"
											r.summary = "delete rule"
											r.operationID = ""
											r.pathPattern = "/api/iam-role-filtering-rules/{ruleID}"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

									elem = origElem
								}

								elem = origElem
							case 's': // Prefix: "s"
								origElem := elem
								if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = "APIIamRolesGet"
										r.summary = "return iam roles in project id"
										r.operationID = ""
										r.pathPattern = "/api/iam-roles"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}

							elem = origElem
						case 'n': // Prefix: "nvitations"
							origElem := elem
							if l := len("nvitations"); len(elem) >= l && elem[0:l] == "nvitations" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "APIInvitationsPost"
									r.summary = "create user invitation"
									r.operationID = ""
									r.pathPattern = "/api/invitations"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					case 'r': // Prefix: "requests"
						origElem := elem
						if l := len("requests"); len(elem) >= l && elem[0:l] == "requests" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "APIRequestsGet"
								r.summary = "return admin request with paging"
								r.operationID = ""
								r.pathPattern = "/api/requests"
								r.args = args
								r.count = 0
								return r, true
							case "POST":
								r.name = "APIRequestsPost"
								r.summary = "post request"
								r.operationID = ""
								r.pathPattern = "/api/requests"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							origElem := elem
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "requestID"
							// Leaf parameter
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "DELETE":
									r.name = "APIRequestsRequestIDDelete"
									r.summary = "delete request"
									r.operationID = ""
									r.pathPattern = "/api/requests/{requestID}"
									r.args = args
									r.count = 1
									return r, true
								case "PATCH":
									r.name = "APIRequestsRequestIDPatch"
									r.summary = "update request"
									r.operationID = ""
									r.pathPattern = "/api/requests/{requestID}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					case 's': // Prefix: "settings"
						origElem := elem
						if l := len("settings"); len(elem) >= l && elem[0:l] == "settings" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PATCH":
								r.name = "APISettingsPatch"
								r.summary = "update settings"
								r.operationID = ""
								r.pathPattern = "/api/settings"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'u': // Prefix: "users"
						origElem := elem
						if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = "APIUsersGet"
								r.summary = "return admin user with paging"
								r.operationID = ""
								r.pathPattern = "/api/users"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							origElem := elem
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "userID"
							// Leaf parameter
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PATCH":
									r.name = "APIUsersUserIDPatch"
									r.summary = "update user"
									r.operationID = ""
									r.pathPattern = "/api/users/{userID}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'u': // Prefix: "uth/google/callback"
					origElem := elem
					if l := len("uth/google/callback"); len(elem) >= l && elem[0:l] == "uth/google/callback" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "AuthGoogleCallbackGet"
							r.summary = "Google callback endpoint"
							r.operationID = ""
							r.pathPattern = "/auth/google/callback"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "health"
				origElem := elem
				if l := len("health"); len(elem) >= l && elem[0:l] == "health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = "HealthGet"
						r.summary = "healthcheck"
						r.operationID = ""
						r.pathPattern = "/health"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'r': // Prefix: "request"
				origElem := elem
				if l := len("request"); len(elem) >= l && elem[0:l] == "request" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "RequestGet"
						r.summary = "return request list page"
						r.operationID = ""
						r.pathPattern = "/request"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '-': // Prefix: "-form"
					origElem := elem
					if l := len("-form"); len(elem) >= l && elem[0:l] == "-form" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "RequestFormGet"
							r.summary = "get request form"
							r.operationID = ""
							r.pathPattern = "/request-form"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "requestID"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "RequestRequestIDGet"
							r.summary = "get request page"
							r.operationID = ""
							r.pathPattern = "/request/{requestID}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "sign"
				origElem := elem
				if l := len("sign"); len(elem) >= l && elem[0:l] == "sign" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "in"
					origElem := elem
					if l := len("in"); len(elem) >= l && elem[0:l] == "in" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "SigninPost"
							r.summary = "sign in"
							r.operationID = ""
							r.pathPattern = "/signin"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'o': // Prefix: "out"
					origElem := elem
					if l := len("out"); len(elem) >= l && elem[0:l] == "out" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "SignoutPost"
							r.summary = "sign out"
							r.operationID = ""
							r.pathPattern = "/signout"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
