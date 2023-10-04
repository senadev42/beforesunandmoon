package main

type StellarObjectData struct {
	Color   []interface{} `json:"color"`
	Catalog []string      `json:"catalog"`
	Dec     []struct {
		Value string `json:"value"`
	} `json:"dec"`
	Escapevelocity         []interface{} `json:"escapevelocity"`
	Name                   []string      `json:"name"`
	Host                   []interface{} `json:"host"`
	Download               []interface{} `json:"download"`
	Masses                 []interface{} `json:"masses"`
	Galactocentricvelocity []interface{} `json:"galactocentricvelocity"`
	Hostoffsetang          []interface{} `json:"hostoffsetang"`
	Xraylink               []interface{} `json:"xraylink"`
	Ra                     []struct {
		Value string `json:"value"`
	} `json:"ra"`
	Propermotionra   []interface{} `json:"propermotionra"`
	Discoverer       []interface{} `json:"discoverer"`
	Propermotiondec  []interface{} `json:"propermotiondec"`
	Instruments      []interface{} `json:"instruments"`
	Photolink        []interface{} `json:"photolink"`
	Redshift         []interface{} `json:"redshift"`
	Boundprobability []interface{} `json:"boundprobability"`
	Claimedtype      []struct {
		Value string `json:"value"`
	} `json:"claimedtype"`
	Radiolink []interface{} `json:"radiolink"`
	Ebv       []struct {
		Value string `json:"value"`
	} `json:"ebv"`
	Hostra         []interface{} `json:"hostra"`
	References     []string      `json:"references"`
	Spectralink    []interface{} `json:"spectralink"`
	Hostoffsetdist []interface{} `json:"hostoffsetdist"`
	Stellarclass   []interface{} `json:"stellarclass"`
	Maxdate        []interface{} `json:"maxdate"`
	Alias          []struct {
		Value string `json:"value"`
	} `json:"alias"`
	Maxappmag    []interface{} `json:"maxappmag"`
	Maxabsmag    []interface{} `json:"maxabsmag"`
	Velocity     []interface{} `json:"velocity"`
	Discoverdate []struct {
		Value string `json:"value"`
	} `json:"discoverdate"`
	Hostdec      []interface{} `json:"hostdec"`
	Spectraltype []interface{} `json:"spectraltype"`
	Lumdist      []interface{} `json:"lumdist"`
}
