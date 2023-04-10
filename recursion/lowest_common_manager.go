package recursion

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	// Write your code here.
	for _, report := range org.DirectReports {
		found := TraverseAndFind(report, reportOne, reportTwo)
		if found == 1 {
			return org
		}
		if found == 2 {
			return GetLowestCommonManager(report, reportOne, reportTwo)
		}
	}
	return org
}

func TraverseAndFind(org, reportOne, reportTwo *OrgChart) int {
	queue := []*OrgChart{org}
	found := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, value := range node.DirectReports {
			queue = append(queue, value)
		}
		if node == reportOne {
			found++
		}
		if node == reportTwo {
			found++
		}
	}
	return found
}
