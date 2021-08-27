package lambertw

var (
	branchPoints = map[int]float64 {
		0:-1, 1:1, 2:-0.333333333333333333e0, 3:0.152777777777777777e0, 4: -0.79629629629629630e-1,
		5: 0.44502314814814814e-1, 6:-0.25984714873603760e-1, 7:0.15635632532333920e-1, 8:-0.96168920242994320e-2,
		9:  0.60145432529561180e-2, 10: -0.38112980348919993e-2, 11:  0.24408779911439826e-2, 12: -0.15769303446867841e-2,
		13: 0.10262633205076071e-2, 14: -0.67206163115613620e-3, 15:  0.44247306181462090e-3, 16: -0.29267722472962746e-3,
		17:  0.19438727605453930e-3, 18: -0.12957426685274883e-3, 19:  0.86650358052081260e-4,
	}
	asymptoticBs = [][]float64{
		{0,-1},{0,1},{0,-1,0.5},{0,1,-3.0/2.0,1.0/3.0},{0,-1,3,-11.0/6.0,0.25},{0,1,-5,35.0/6.0,-25.0/12.0,0.2},
	}
	q = []float64{
		-1,
		+1,
		-0.333333333333333333,
		+0.152777777777777778,
		-0.0796296296296296296,
		+0.0445023148148148148,
		-0.0259847148736037625,
		+0.0156356325323339212,
		-0.00961689202429943171,
		+0.00601454325295611786,
		-0.00381129803489199923,
		+0.00244087799114398267,
		-0.00157693034468678425,
		+0.00102626332050760715,
		-0.000672061631156136204,
		+0.000442473061814620910,
		-0.000292677224729627445,
		+0.000194387276054539318,
		-0.000129574266852748819,
		+0.0000866503580520812717,
		-0.0000581136075044138168,
	}
)