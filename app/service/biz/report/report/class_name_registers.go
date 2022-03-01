/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package report

const (
	Predicate_report_accountReportPeer           = "report_accountReportPeer"
	Predicate_report_accountReportProfilePhoto   = "report_accountReportProfilePhoto"
	Predicate_report_messagesReportSpam          = "report_messagesReportSpam"
	Predicate_report_messagesReport              = "report_messagesReport"
	Predicate_report_messagesReportEncryptedSpam = "report_messagesReportEncryptedSpam"
	Predicate_report_channelsReportSpam          = "report_channelsReportSpam"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_report_accountReportPeer: {
		0: 1976979630, // 0x75d650ae

	},
	Predicate_report_accountReportProfilePhoto: {
		0: -1206920954, // 0xb80fd906

	},
	Predicate_report_messagesReportSpam: {
		0: -2120170998, // 0x81a0c20a

	},
	Predicate_report_messagesReport: {
		0: -1299590501, // 0xb289d29b

	},
	Predicate_report_messagesReportEncryptedSpam: {
		0: 762034535, // 0x2d6bb967

	},
	Predicate_report_channelsReportSpam: {
		0: 2010319160, // 0x77d30938

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1976979630:  Predicate_report_accountReportPeer,           // 0x75d650ae
	-1206920954: Predicate_report_accountReportProfilePhoto,   // 0xb80fd906
	-2120170998: Predicate_report_messagesReportSpam,          // 0x81a0c20a
	-1299590501: Predicate_report_messagesReport,              // 0xb289d29b
	762034535:   Predicate_report_messagesReportEncryptedSpam, // 0x2d6bb967
	2010319160:  Predicate_report_channelsReportSpam,          // 0x77d30938

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
