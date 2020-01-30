// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

func init() {
	// Merge language subs with the default one.
	// TODO: Find better way so all langs are merged automatically and better
	// tested.
	for _, sub := range []*map[rune]string{
		&deSub, &enSub, &esSub, &fiSub, &frSub, &grSub, &kkSub, &nlSub, &plSub, &svSub, &trSub,
	} {
		for key, value := range defaultSub {
			(*sub)[key] = value
		}
	}
}

var defaultSub = map[rune]string{
	'"':  "",
	'\'': "",
	'’':  "",
	'‒':  "-", // figure dash
	'–':  "-", // en dash
	'—':  "-", // em dash
	'―':  "-", // horizontal bar
}

var deSub = map[rune]string{
	'&': "und",
	'@': "at",
	'ä': "ae",
	'Ä': "Ae",
	'ö': "oe",
	'Ö': "Oe",
	'ü': "ue",
	'Ü': "Ue",
	'é': "e",
	'ß': "ss",
}

var enSub = map[rune]string{
	'&': "and",
	'@': "at",
}

var esSub = map[rune]string{
	'&': "y",
	'@': "en",
	'á': "a",
	'é': "e",
	'í': "i",
	'ó': "o",
	'ú': "u",
	'ü': "u",
	'ñ': "n",
	'Á': "A",
	'É': "E",
	'Í': "I",
	'Ó': "O",
	'Ú': "U",
	'Ü': "U",
	'Ñ': "N",
}

var fiSub = map[rune]string{
	'&': "ja",
	'@': "at",
}

var frSub = map[rune]string{
	'&': "et",
	'@': "arobase",
	'à': "a",
	'â': "a",
	'æ': "ae",
	'ç': "c",
	'é': "e",
	'è': "e",
	'ê': "e",
	'ë': "e",
	'î': "i",
	'ï': "i",
	'ô': "o",
	'œ': "oe",
	'ù': "u",
	'û': "u",
	'ü': "u",
	'ÿ': "y",
	'À': "A",
	'Â': "A",
	'Æ': "Ae",
	'Ç': "C",
	'É': "E",
	'È': "E",
	'Ê': "E",
	'Ë': "E",
	'Î': "I",
	'Ï': "I",
	'Ô': "O",
	'Œ': "Oe",
	'Ù': "U",
	'Û': "U",
	'Ü': "U",
	'Ÿ': "Y",
}

var grSub = map[rune]string{
	'&': "kai",
	'η': "i",
	'ή': "i",
	'Η': "i",
	'ι': "i",
	'ί': "i",
	'ϊ': "i",
	'Ι': "i",
	'χ': "x",
	'Χ': "x",
	'ω': "w",
	'ώ': "w",
	'Ω': "w",
	'ϋ': "u",
}

var kkSub = map[rune]string{
	'&': "jane",
	'ә': "a",
	'ғ': "g",
	'қ': "q",
	'ң': "n",
	'ө': "o",
	'ұ': "u",
	'Ә': "A",
	'Ғ': "G",
	'Қ': "Q",
	'Ң': "N",
	'Ө': "O",
	'Ұ': "U",
}

var nlSub = map[rune]string{
	'&': "en",
	'@': "at",
}

var plSub = map[rune]string{
	'&': "i",
	'@': "na",
}

var svSub = map[rune]string{
	'&': "och",
	'@': "snabel a",
}

var trSub = map[rune]string{
	'&': "ve",
	'@': "et",
	'ş': "s",
	'Ş': "S",
	'ü': "u",
	'Ü': "U",
	'ö': "o",
	'Ö': "O",
	'İ': "I",
	'ı': "i",
	'ğ': "g",
	'Ğ': "G",
	'ç': "c",
	'Ç': "C",
}
