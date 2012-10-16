package models

import (
	"github.com/gosexy/sugar"
	"github.com/astrata/tango/app"
	"github.com/astrata/tango/body"
	"labix.org/v2/mgo/bson"
	"labix.org/v2/mgo"
	"persistent"
	"strings"
)

type Country struct {
	Database *mgo.Database
	Collection *mgo.Collection
}

func init() {
	app.Register("Country", &Country{})
	app.Route("/country", app.App("Country"))
}

func (self *Country) StartUp() {

	self.Database		= persistent.Database("default")

	self.Collection = self.Database.C("countries")

	self.Collection.DropCollection()

	self.Collection.Insert(
		bson.M{"code": "AF", "name": "Afghanistan"},
		bson.M{"code": "AX", "name": "Åland Islands"},
		bson.M{"code": "AL", "name": "Albania"},
		bson.M{"code": "DZ", "name": "Algeria"},
		bson.M{"code": "AS", "name": "American Samoa"},
		bson.M{"code": "AD", "name": "Andorra"},
		bson.M{"code": "AO", "name": "Angola"},
		bson.M{"code": "AI", "name": "Anguilla"},
		bson.M{"code": "AQ", "name": "Antarctica"},
		bson.M{"code": "AG", "name": "Antigua and Barbuda"},
		bson.M{"code": "AR", "name": "Argentina"},
		bson.M{"code": "AM", "name": "Armenia"},
		bson.M{"code": "AW", "name": "Aruba"},
		bson.M{"code": "AU", "name": "Australia"},
		bson.M{"code": "AT", "name": "Austria"},
		bson.M{"code": "AZ", "name": "Azerbaijan"},
		bson.M{"code": "BS", "name": "Bahamas"},
		bson.M{"code": "BH", "name": "Bahrain"},
		bson.M{"code": "BD", "name": "Bangladesh"},
		bson.M{"code": "BB", "name": "Barbados"},
		bson.M{"code": "BY", "name": "Belarus"},
		bson.M{"code": "BE", "name": "Belgium"},
		bson.M{"code": "BZ", "name": "Belize"},
		bson.M{"code": "BJ", "name": "Benin"},
		bson.M{"code": "BM", "name": "Bermuda"},
		bson.M{"code": "BT", "name": "Bhutan"},
		bson.M{"code": "BO", "name": "Bolivia"},
		bson.M{"code": "BA", "name": "Bosnia and Herzegovina"},
		bson.M{"code": "BW", "name": "Botswana"},
		bson.M{"code": "BV", "name": "Bouvet Island"},
		bson.M{"code": "BR", "name": "Brazil"},
		bson.M{"code": "IO", "name": "British Indian Ocean Territory"},
		bson.M{"code": "BN", "name": "Brunei Darussalam"},
		bson.M{"code": "BG", "name": "Bulgaria"},
		bson.M{"code": "BF", "name": "Burkina Faso"},
		bson.M{"code": "BI", "name": "Burundi"},
		bson.M{"code": "KH", "name": "Cambodia"},
		bson.M{"code": "CM", "name": "Cameroon"},
		bson.M{"code": "CA", "name": "Canada"},
		bson.M{"code": "CV", "name": "Cape Verde"},
		bson.M{"code": "KY", "name": "Cayman Islands"},
		bson.M{"code": "CF", "name": "Central African Republic"},
		bson.M{"code": "TD", "name": "Chad"},
		bson.M{"code": "CL", "name": "Chile"},
		bson.M{"code": "CN", "name": "China"},
		bson.M{"code": "CX", "name": "Christmas Island"},
		bson.M{"code": "CC", "name": "Cocos (Keeling) Islands"},
		bson.M{"code": "CO", "name": "Colombia"},
		bson.M{"code": "KM", "name": "Comoros"},
		bson.M{"code": "CG", "name": "Congo"},
		bson.M{"code": "CD", "name": "Congo, The Democratic Republic of The"},
		bson.M{"code": "CK", "name": "Cook Islands"},
		bson.M{"code": "CR", "name": "Costa Rica"},
		bson.M{"code": "CI", "name": "Cote D'ivoire"},
		bson.M{"code": "HR", "name": "Croatia"},
		bson.M{"code": "CU", "name": "Cuba"},
		bson.M{"code": "CY", "name": "Cyprus"},
		bson.M{"code": "CZ", "name": "Czech Republic"},
		bson.M{"code": "DK", "name": "Denmark"},
		bson.M{"code": "DJ", "name": "Djibouti"},
		bson.M{"code": "DM", "name": "Dominica"},
		bson.M{"code": "DO", "name": "Dominican Republic"},
		bson.M{"code": "EC", "name": "Ecuador"},
		bson.M{"code": "EG", "name": "Egypt"},
		bson.M{"code": "SV", "name": "El Salvador"},
		bson.M{"code": "GQ", "name": "Equatorial Guinea"},
		bson.M{"code": "ER", "name": "Eritrea"},
		bson.M{"code": "EE", "name": "Estonia"},
		bson.M{"code": "ET", "name": "Ethiopia"},
		bson.M{"code": "FK", "name": "Falkland Islands (Malvinas)"},
		bson.M{"code": "FO", "name": "Faroe Islands"},
		bson.M{"code": "FJ", "name": "Fiji"},
		bson.M{"code": "FI", "name": "Finland"},
		bson.M{"code": "FR", "name": "France"},
		bson.M{"code": "GF", "name": "French Guiana"},
		bson.M{"code": "PF", "name": "French Polynesia"},
		bson.M{"code": "TF", "name": "French Southern Territories"},
		bson.M{"code": "GA", "name": "Gabon"},
		bson.M{"code": "GM", "name": "Gambia"},
		bson.M{"code": "GE", "name": "Georgia"},
		bson.M{"code": "DE", "name": "Germany"},
		bson.M{"code": "GH", "name": "Ghana"},
		bson.M{"code": "GI", "name": "Gibraltar"},
		bson.M{"code": "GR", "name": "Greece"},
		bson.M{"code": "GL", "name": "Greenland"},
		bson.M{"code": "GD", "name": "Grenada"},
		bson.M{"code": "GP", "name": "Guadeloupe"},
		bson.M{"code": "GU", "name": "Guam"},
		bson.M{"code": "GT", "name": "Guatemala"},
		bson.M{"code": "GG", "name": "Guernsey"},
		bson.M{"code": "GN", "name": "Guinea"},
		bson.M{"code": "GW", "name": "Guinea-bissau"},
		bson.M{"code": "GY", "name": "Guyana"},
		bson.M{"code": "HT", "name": "Haiti"},
		bson.M{"code": "HM", "name": "Heard Island and Mcdonald Islands"},
		bson.M{"code": "VA", "name": "Holy See (Vatican City State)"},
		bson.M{"code": "HN", "name": "Honduras"},
		bson.M{"code": "HK", "name": "Hong Kong"},
		bson.M{"code": "HU", "name": "Hungary"},
		bson.M{"code": "IS", "name": "Iceland"},
		bson.M{"code": "IN", "name": "India"},
		bson.M{"code": "ID", "name": "Indonesia"},
		bson.M{"code": "IR", "name": "Iran, Islamic Republic of"},
		bson.M{"code": "IQ", "name": "Iraq"},
		bson.M{"code": "IE", "name": "Ireland"},
		bson.M{"code": "IM", "name": "Isle of Man"},
		bson.M{"code": "IL", "name": "Israel"},
		bson.M{"code": "IT", "name": "Italy"},
		bson.M{"code": "JM", "name": "Jamaica"},
		bson.M{"code": "JP", "name": "Japan"},
		bson.M{"code": "JE", "name": "Jersey"},
		bson.M{"code": "JO", "name": "Jordan"},
		bson.M{"code": "KZ", "name": "Kazakhstan"},
		bson.M{"code": "KE", "name": "Kenya"},
		bson.M{"code": "KI", "name": "Kiribati"},
		bson.M{"code": "KP", "name": "Korea, Democratic People's Republic of"},
		bson.M{"code": "KR", "name": "Korea, Republic of"},
		bson.M{"code": "KW", "name": "Kuwait"},
		bson.M{"code": "KG", "name": "Kyrgyzstan"},
		bson.M{"code": "LA", "name": "Lao People's Democratic Republic"},
		bson.M{"code": "LV", "name": "Latvia"},
		bson.M{"code": "LB", "name": "Lebanon"},
		bson.M{"code": "LS", "name": "Lesotho"},
		bson.M{"code": "LR", "name": "Liberia"},
		bson.M{"code": "LY", "name": "Libyan Arab Jamahiriya"},
		bson.M{"code": "LI", "name": "Liechtenstein"},
		bson.M{"code": "LT", "name": "Lithuania"},
		bson.M{"code": "LU", "name": "Luxembourg"},
		bson.M{"code": "MO", "name": "Macao"},
		bson.M{"code": "MK", "name": "Macedonia, The Former Yugoslav Republic of"},
		bson.M{"code": "MG", "name": "Madagascar"},
		bson.M{"code": "MW", "name": "Malawi"},
		bson.M{"code": "MY", "name": "Malaysia"},
		bson.M{"code": "MV", "name": "Maldives"},
		bson.M{"code": "ML", "name": "Mali"},
		bson.M{"code": "MT", "name": "Malta"},
		bson.M{"code": "MH", "name": "Marshall Islands"},
		bson.M{"code": "MQ", "name": "Martinique"},
		bson.M{"code": "MR", "name": "Mauritania"},
		bson.M{"code": "MU", "name": "Mauritius"},
		bson.M{"code": "YT", "name": "Mayotte"},
		bson.M{"code": "MX", "name": "Mexico"},
		bson.M{"code": "FM", "name": "Micronesia, Federated States of"},
		bson.M{"code": "MD", "name": "Moldova, Republic of"},
		bson.M{"code": "MC", "name": "Monaco"},
		bson.M{"code": "MN", "name": "Mongolia"},
		bson.M{"code": "ME", "name": "Montenegro"},
		bson.M{"code": "MS", "name": "Montserrat"},
		bson.M{"code": "MA", "name": "Morocco"},
		bson.M{"code": "MZ", "name": "Mozambique"},
		bson.M{"code": "MM", "name": "Myanmar"},
		bson.M{"code": "NA", "name": "Namibia"},
		bson.M{"code": "NR", "name": "Nauru"},
		bson.M{"code": "NP", "name": "Nepal"},
		bson.M{"code": "NL", "name": "Netherlands"},
		bson.M{"code": "AN", "name": "Netherlands Antilles"},
		bson.M{"code": "NC", "name": "New Caledonia"},
		bson.M{"code": "NZ", "name": "New Zealand"},
		bson.M{"code": "NI", "name": "Nicaragua"},
		bson.M{"code": "NE", "name": "Niger"},
		bson.M{"code": "NG", "name": "Nigeria"},
		bson.M{"code": "NU", "name": "Niue"},
		bson.M{"code": "NF", "name": "Norfolk Island"},
		bson.M{"code": "MP", "name": "Northern Mariana Islands"},
		bson.M{"code": "NO", "name": "Norway"},
		bson.M{"code": "OM", "name": "Oman"},
		bson.M{"code": "PK", "name": "Pakistan"},
		bson.M{"code": "PW", "name": "Palau"},
		bson.M{"code": "PS", "name": "Palestinian Territory, Occupied"},
		bson.M{"code": "PA", "name": "Panama"},
		bson.M{"code": "PG", "name": "Papua New Guinea"},
		bson.M{"code": "PY", "name": "Paraguay"},
		bson.M{"code": "PE", "name": "Peru"},
		bson.M{"code": "PH", "name": "Philippines"},
		bson.M{"code": "PN", "name": "Pitcairn"},
		bson.M{"code": "PL", "name": "Poland"},
		bson.M{"code": "PT", "name": "Portugal"},
		bson.M{"code": "PR", "name": "Puerto Rico"},
		bson.M{"code": "QA", "name": "Qatar"},
		bson.M{"code": "RE", "name": "Reunion"},
		bson.M{"code": "RO", "name": "Romania"},
		bson.M{"code": "RU", "name": "Russian Federation"},
		bson.M{"code": "RW", "name": "Rwanda"},
		bson.M{"code": "SH", "name": "Saint Helena"},
		bson.M{"code": "KN", "name": "Saint Kitts and Nevis"},
		bson.M{"code": "LC", "name": "Saint Lucia"},
		bson.M{"code": "PM", "name": "Saint Pierre and Miquelon"},
		bson.M{"code": "VC", "name": "Saint Vincent and The Grenadines"},
		bson.M{"code": "WS", "name": "Samoa"},
		bson.M{"code": "SM", "name": "San Marino"},
		bson.M{"code": "ST", "name": "Sao Tome and Principe"},
		bson.M{"code": "SA", "name": "Saudi Arabia"},
		bson.M{"code": "SN", "name": "Senegal"},
		bson.M{"code": "RS", "name": "Serbia"},
		bson.M{"code": "SC", "name": "Seychelles"},
		bson.M{"code": "SL", "name": "Sierra Leone"},
		bson.M{"code": "SG", "name": "Singapore"},
		bson.M{"code": "SK", "name": "Slovakia"},
		bson.M{"code": "SI", "name": "Slovenia"},
		bson.M{"code": "SB", "name": "Solomon Islands"},
		bson.M{"code": "SO", "name": "Somalia"},
		bson.M{"code": "ZA", "name": "South Africa"},
		bson.M{"code": "GS", "name": "South Georgia and The South Sandwich Islands"},
		bson.M{"code": "ES", "name": "Spain"},
		bson.M{"code": "LK", "name": "Sri Lanka"},
		bson.M{"code": "SD", "name": "Sudan"},
		bson.M{"code": "SR", "name": "Suriname"},
		bson.M{"code": "SJ", "name": "Svalbard and Jan Mayen"},
		bson.M{"code": "SZ", "name": "Swaziland"},
		bson.M{"code": "SE", "name": "Sweden"},
		bson.M{"code": "CH", "name": "Switzerland"},
		bson.M{"code": "SY", "name": "Syrian Arab Republic"},
		bson.M{"code": "TW", "name": "Taiwan, Province of China"},
		bson.M{"code": "TJ", "name": "Tajikistan"},
		bson.M{"code": "TZ", "name": "Tanzania, United Republic of"},
		bson.M{"code": "TH", "name": "Thailand"},
		bson.M{"code": "TL", "name": "Timor-leste"},
		bson.M{"code": "TG", "name": "Togo"},
		bson.M{"code": "TK", "name": "Tokelau"},
		bson.M{"code": "TO", "name": "Tonga"},
		bson.M{"code": "TT", "name": "Trinidad and Tobago"},
		bson.M{"code": "TN", "name": "Tunisia"},
		bson.M{"code": "TR", "name": "Turkey"},
		bson.M{"code": "TM", "name": "Turkmenistan"},
		bson.M{"code": "TC", "name": "Turks and Caicos Islands"},
		bson.M{"code": "TV", "name": "Tuvalu"},
		bson.M{"code": "UG", "name": "Uganda"},
		bson.M{"code": "UA", "name": "Ukraine"},
		bson.M{"code": "AE", "name": "United Arab Emirates"},
		bson.M{"code": "GB", "name": "United Kingdom"},
		bson.M{"code": "US", "name": "United States"},
		bson.M{"code": "UM", "name": "United States Minor Outlying Islands"},
		bson.M{"code": "UY", "name": "Uruguay"},
		bson.M{"code": "UZ", "name": "Uzbekistan"},
		bson.M{"code": "VU", "name": "Vanuatu"},
		bson.M{"code": "VE", "name": "Venezuela"},
		bson.M{"code": "VN", "name": "Viet Nam"},
		bson.M{"code": "VG", "name": "Virgin Islands, British"},
		bson.M{"code": "VI", "name": "Virgin Islands, U.S."},
		bson.M{"code": "WF", "name": "Wallis and Futuna"},
		bson.M{"code": "EH", "name": "Western Sahara"},
		bson.M{"code": "YE", "name": "Yemen"},
		bson.M{"code": "ZM", "name": "Zambia"},
		bson.M{"code": "ZW", "name": "Zimbabwe"},
	)
}

type CountryValue struct {
	Code string
	Name string
}

func (self *Country) Find(code string) body.Body {

	response := body.Json()

	country := &CountryValue{}

	err := self.Collection.Find(
		bson.M {
			"code": strings.ToUpper(code),
		},
	).One(country)

	if err != nil {
		response.Set(
			sugar.Tuple {
				"error": "Not found",
			},
		)
	} else {
		response.Set(
			sugar.Tuple {
				"success": "Found",
				"data": sugar.Tuple{
					"code": country.Code,
					"name": country.Name,
				},
			},
		)
	}

	return response
}

