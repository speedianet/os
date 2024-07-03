package infraHelper

import "strings"

func DnsLookup(recordName string, recordType *string) ([]string, error) {
	resourceRecords := []string{}

	recordTypeStr := "A"
	if recordType != nil {
		recordTypeStr = *recordType
	}

	digCmd := "dig +short " + recordTypeStr + " " + recordName

	rawRecords, err := RunCmdWithSubShell(digCmd + " @dns.google")
	if err != nil || rawRecords == "" {
		rawRecords, err = RunCmdWithSubShell(digCmd + " @security-filter-dns.cleanbrowsing.org")
		if err != nil {
			return resourceRecords, err
		}
	}

	if rawRecords == "" {
		return resourceRecords, nil
	}

	rawRecordsParts := strings.Split(rawRecords, "\n")
	for _, rawRecord := range rawRecordsParts {
		resourceRecords = append(resourceRecords, strings.TrimSpace(rawRecord))
	}

	return resourceRecords, nil
}
