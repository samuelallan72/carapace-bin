package dagger

import "github.com/rsteube/carapace"

// ActionLicenses completes licenses
//
//	0BSD (BSD Zero Clause License)
//	AAL (Attribution Assurance License)
func ActionLicenses() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0BSD", "BSD Zero Clause License",
		"AAL", "Attribution Assurance License",
		"Abstyles", "Abstyles License",
		"AdaCore-doc", "AdaCore Doc License",
		"Adobe-2006", "Adobe Systems Incorporated Source Code License Agreement",
		"Adobe-Display-PostScript", "Adobe Display PostScript License",
		"Adobe-Glyph", "Adobe Glyph List License",
		"Adobe-Utopia", "Adobe Utopia Font License",
		"ADSL", "Amazon Digital Services License",
		"AFL-1.1", "Academic Free License v1.1",
		"AFL-1.2", "Academic Free License v1.2",
		"AFL-2.0", "Academic Free License v2.0",
		"AFL-2.1", "Academic Free License v2.1",
		"AFL-3.0", "Academic Free License v3.0",
		"Afmparse", "Afmparse License",
		"AGPL-1.0", "Affero General Public License v1.0",
		"AGPL-1.0-only", "Affero General Public License v1.0 only",
		"AGPL-1.0-or-later", "Affero General Public License v1.0 or later",
		"AGPL-3.0", "GNU Affero General Public License v3.0",
		"AGPL-3.0-only", "GNU Affero General Public License v3.0 only",
		"AGPL-3.0-or-later", "GNU Affero General Public License v3.0 or later",
		"Aladdin", "Aladdin Free Public License",
		"AMDPLPA", "AMD's plpa_map.c License",
		"AML", "Apple MIT License",
		"AML-glslang", "AML glslang variant License",
		"AMPAS", "Academy of Motion Picture Arts and Sciences BSD",
		"ANTLR-PD", "ANTLR Software Rights Notice",
		"ANTLR-PD-fallback", "ANTLR Software Rights Notice with license fallback",
		"Apache-1.0", "Apache License 1.0",
		"Apache-1.1", "Apache License 1.1",
		"Apache-2.0", "Apache License 2.0",
		"APAFML", "Adobe Postscript AFM License",
		"APL-1.0", "Adaptive Public License 1.0",
		"App-s2p", "App::s2p License",
		"APSL-1.0", "Apple Public Source License 1.0",
		"APSL-1.1", "Apple Public Source License 1.1",
		"APSL-1.2", "Apple Public Source License 1.2",
		"APSL-2.0", "Apple Public Source License 2.0",
		"Arphic-1999", "Arphic Public License",
		"Artistic-1.0", "Artistic License 1.0",
		"Artistic-1.0-cl8", "Artistic License 1.0 w/clause 8",
		"Artistic-1.0-Perl", "Artistic License 1.0 (Perl)",
		"Artistic-2.0", "Artistic License 2.0",
		"ASWF-Digital-Assets-1.0", "ASWF Digital Assets License version 1.0",
		"ASWF-Digital-Assets-1.1", "ASWF Digital Assets License 1.1",
		"Baekmuk", "Baekmuk License",
		"Bahyph", "Bahyph License",
		"Barr", "Barr License",
		"bcrypt-Solar-Designer", "bcrypt Solar Designer License",
		"Beerware", "Beerware License",
		"Bitstream-Charter", "Bitstream Charter Font License",
		"Bitstream-Vera", "Bitstream Vera Font License",
		"BitTorrent-1.0", "BitTorrent Open Source License v1.0",
		"BitTorrent-1.1", "BitTorrent Open Source License v1.1",
		"blessing", "SQLite Blessing",
		"BlueOak-1.0.0", "Blue Oak Model License 1.0.0",
		"Boehm-GC", "Boehm-Demers-Weiser GC License",
		"Borceux", "Borceux license",
		"Brian-Gladman-2-Clause", "Brian Gladman 2-Clause License",
		"Brian-Gladman-3-Clause", "Brian Gladman 3-Clause License",
		"BSD-1-Clause", "BSD 1-Clause License",
		"BSD-2-Clause", "BSD 2-Clause \"Simplified\" License",
		"BSD-2-Clause-Darwin", "BSD 2-Clause - Ian Darwin variant",
		"BSD-2-Clause-FreeBSD", "BSD 2-Clause FreeBSD License",
		"BSD-2-Clause-NetBSD", "BSD 2-Clause NetBSD License",
		"BSD-2-Clause-Patent", "BSD-2-Clause Plus Patent License",
		"BSD-2-Clause-Views", "BSD 2-Clause with views sentence",
		"BSD-3-Clause-acpica", "BSD 3-Clause acpica variant",
		"BSD-3-Clause-Attribution", "BSD with attribution",
		"BSD-3-Clause", "BSD 3-Clause \"New\" or \"Revised\" License",
		"BSD-3-Clause-Clear", "BSD 3-Clause Clear License",
		"BSD-3-Clause-flex", "BSD 3-Clause Flex variant",
		"BSD-3-Clause-HP", "Hewlett-Packard BSD variant license",
		"BSD-3-Clause-LBNL", "Lawrence Berkeley National Labs BSD variant license",
		"BSD-3-Clause-Modification", "BSD 3-Clause Modification",
		"BSD-3-Clause-No-Military-License", "BSD 3-Clause No Military License",
		"BSD-3-Clause-No-Nuclear-License-2014", "BSD 3-Clause No Nuclear License 2014",
		"BSD-3-Clause-No-Nuclear-License", "BSD 3-Clause No Nuclear License",
		"BSD-3-Clause-No-Nuclear-Warranty", "BSD 3-Clause No Nuclear Warranty",
		"BSD-3-Clause-Open-MPI", "BSD 3-Clause Open MPI variant",
		"BSD-3-Clause-Sun", "BSD 3-Clause Sun Microsystems",
		"BSD-4.3RENO", "BSD 4.3 RENO License",
		"BSD-4.3TAHOE", "BSD 4.3 TAHOE License",
		"BSD-4-Clause", "BSD 4-Clause \"Original\" or \"Old\" License",
		"BSD-4-Clause-Shortened", "BSD 4 Clause Shortened",
		"BSD-4-Clause-UC", "BSD-4-Clause (University of California-Specific)",
		"BSD-Advertising-Acknowledgement", "BSD Advertising Acknowledgement License",
		"BSD-Attribution-HPND-disclaimer", "BSD with Attribution and HPND disclaimer",
		"BSD-Inferno-Nettverk", "BSD-Inferno-Nettverk",
		"BSD-Protection", "BSD Protection License",
		"BSD-Source-beginning-file", "BSD Source Code Attribution - beginning of file variant",
		"BSD-Source-Code", "BSD Source Code Attribution",
		"BSD-Systemics", "Systemics BSD variant license",
		"BSD-Systemics-W3Works", "Systemics W3Works BSD variant license",
		"BSL-1.0", "Boost Software License 1.0",
		"BUSL-1.1", "Business Source License 1.1",
		"bzip2-1.0.5", "bzip2 and libbzip2 License v1.0.5",
		"bzip2-1.0.6", "bzip2 and libbzip2 License v1.0.6",
		"CAL-1.0-Combined-Work-Exception", "Cryptographic Autonomy License 1.0 (Combined Work Exception)",
		"CAL-1.0", "Cryptographic Autonomy License 1.0",
		"Caldera", "Caldera License",
		"Caldera-no-preamble", "Caldera License (without preamble)",
		"CATOSL-1.1", "Computer Associates Trusted Open Source License 1.1",
		"CC0-1.0", "Creative Commons Zero v1.0 Universal",
		"CC-BY-1.0", "Creative Commons Attribution 1.0 Generic",
		"CC-BY-2.0", "Creative Commons Attribution 2.0 Generic",
		"CC-BY-2.5-AU", "Creative Commons Attribution 2.5 Australia",
		"CC-BY-2.5", "Creative Commons Attribution 2.5 Generic",
		"CC-BY-3.0-AT", "Creative Commons Attribution 3.0 Austria",
		"CC-BY-3.0-AU", "Creative Commons Attribution 3.0 Australia",
		"CC-BY-3.0", "Creative Commons Attribution 3.0 Unported",
		"CC-BY-3.0-DE", "Creative Commons Attribution 3.0 Germany",
		"CC-BY-3.0-IGO", "Creative Commons Attribution 3.0 IGO",
		"CC-BY-3.0-NL", "Creative Commons Attribution 3.0 Netherlands",
		"CC-BY-3.0-US", "Creative Commons Attribution 3.0 United States",
		"CC-BY-4.0", "Creative Commons Attribution 4.0 International",
		"CC-BY-NC-1.0", "Creative Commons Attribution Non Commercial 1.0 Generic",
		"CC-BY-NC-2.0", "Creative Commons Attribution Non Commercial 2.0 Generic",
		"CC-BY-NC-2.5", "Creative Commons Attribution Non Commercial 2.5 Generic",
		"CC-BY-NC-3.0", "Creative Commons Attribution Non Commercial 3.0 Unported",
		"CC-BY-NC-3.0-DE", "Creative Commons Attribution Non Commercial 3.0 Germany",
		"CC-BY-NC-4.0", "Creative Commons Attribution Non Commercial 4.0 International",
		"CC-BY-NC-ND-1.0", "Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic",
		"CC-BY-NC-ND-2.0", "Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic",
		"CC-BY-NC-ND-2.5", "Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic",
		"CC-BY-NC-ND-3.0", "Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported",
		"CC-BY-NC-ND-3.0-DE", "Creative Commons Attribution Non Commercial No Derivatives 3.0 Germany",
		"CC-BY-NC-ND-3.0-IGO", "Creative Commons Attribution Non Commercial No Derivatives 3.0 IGO",
		"CC-BY-NC-ND-4.0", "Creative Commons Attribution Non Commercial No Derivatives 4.0 International",
		"CC-BY-NC-SA-1.0", "Creative Commons Attribution Non Commercial Share Alike 1.0 Generic",
		"CC-BY-NC-SA-2.0", "Creative Commons Attribution Non Commercial Share Alike 2.0 Generic",
		"CC-BY-NC-SA-2.0-DE", "Creative Commons Attribution Non Commercial Share Alike 2.0 Germany",
		"CC-BY-NC-SA-2.0-FR", "Creative Commons Attribution-NonCommercial-ShareAlike 2.0 France",
		"CC-BY-NC-SA-2.0-UK", "Creative Commons Attribution Non Commercial Share Alike 2.0 England and Wales",
		"CC-BY-NC-SA-2.5", "Creative Commons Attribution Non Commercial Share Alike 2.5 Generic",
		"CC-BY-NC-SA-3.0", "Creative Commons Attribution Non Commercial Share Alike 3.0 Unported",
		"CC-BY-NC-SA-3.0-DE", "Creative Commons Attribution Non Commercial Share Alike 3.0 Germany",
		"CC-BY-NC-SA-3.0-IGO", "Creative Commons Attribution Non Commercial Share Alike 3.0 IGO",
		"CC-BY-NC-SA-4.0", "Creative Commons Attribution Non Commercial Share Alike 4.0 International",
		"CC-BY-ND-1.0", "Creative Commons Attribution No Derivatives 1.0 Generic",
		"CC-BY-ND-2.0", "Creative Commons Attribution No Derivatives 2.0 Generic",
		"CC-BY-ND-2.5", "Creative Commons Attribution No Derivatives 2.5 Generic",
		"CC-BY-ND-3.0", "Creative Commons Attribution No Derivatives 3.0 Unported",
		"CC-BY-ND-3.0-DE", "Creative Commons Attribution No Derivatives 3.0 Germany",
		"CC-BY-ND-4.0", "Creative Commons Attribution No Derivatives 4.0 International",
		"CC-BY-SA-1.0", "Creative Commons Attribution Share Alike 1.0 Generic",
		"CC-BY-SA-2.0", "Creative Commons Attribution Share Alike 2.0 Generic",
		"CC-BY-SA-2.0-UK", "Creative Commons Attribution Share Alike 2.0 England and Wales",
		"CC-BY-SA-2.1-JP", "Creative Commons Attribution Share Alike 2.1 Japan",
		"CC-BY-SA-2.5", "Creative Commons Attribution Share Alike 2.5 Generic",
		"CC-BY-SA-3.0-AT", "Creative Commons Attribution Share Alike 3.0 Austria",
		"CC-BY-SA-3.0", "Creative Commons Attribution Share Alike 3.0 Unported",
		"CC-BY-SA-3.0-DE", "Creative Commons Attribution Share Alike 3.0 Germany",
		"CC-BY-SA-3.0-IGO", "Creative Commons Attribution-ShareAlike 3.0 IGO",
		"CC-BY-SA-4.0", "Creative Commons Attribution Share Alike 4.0 International",
		"CC-PDDC", "Creative Commons Public Domain Dedication and Certification",
		"CDDL-1.0", "Common Development and Distribution License 1.0",
		"CDDL-1.1", "Common Development and Distribution License 1.1",
		"CDL-1.0", "Common Documentation License 1.0",
		"CDLA-Permissive-1.0", "Community Data License Agreement Permissive 1.0",
		"CDLA-Permissive-2.0", "Community Data License Agreement Permissive 2.0",
		"CDLA-Sharing-1.0", "Community Data License Agreement Sharing 1.0",
		"CECILL-1.0", "CeCILL Free Software License Agreement v1.0",
		"CECILL-1.1", "CeCILL Free Software License Agreement v1.1",
		"CECILL-2.0", "CeCILL Free Software License Agreement v2.0",
		"CECILL-2.1", "CeCILL Free Software License Agreement v2.1",
		"CECILL-B", "CeCILL-B Free Software License Agreement",
		"CECILL-C", "CeCILL-C Free Software License Agreement",
		"CERN-OHL-1.1", "CERN Open Hardware Licence v1.1",
		"CERN-OHL-1.2", "CERN Open Hardware Licence v1.2",
		"CERN-OHL-P-2.0", "CERN Open Hardware Licence Version 2 - Permissive",
		"CERN-OHL-S-2.0", "CERN Open Hardware Licence Version 2 - Strongly Reciprocal",
		"CERN-OHL-W-2.0", "CERN Open Hardware Licence Version 2 - Weakly Reciprocal",
		"CFITSIO", "CFITSIO License",
		"check-cvs", "check-cvs License",
		"checkmk", "Checkmk License",
		"ClArtistic", "Clarified Artistic License",
		"Clips", "Clips License",
		"CMU-Mach", "CMU Mach License",
		"CMU-Mach-nodoc", "CMU    Mach - no notices-in-documentation variant",
		"CNRI-Jython", "CNRI Jython License",
		"CNRI-Python", "CNRI Python License",
		"CNRI-Python-GPL-Compatible", "CNRI Python Open Source GPL Compatible License Agreement",
		"COIL-1.0", "Copyfree Open Innovation License",
		"Community-Spec-1.0", "Community Specification License 1.0",
		"Condor-1.1", "Condor Public License v1.1",
		"copyleft-next-0.3.0", "copyleft-next 0.3.0",
		"copyleft-next-0.3.1", "copyleft-next 0.3.1",
		"Cornell-Lossless-JPEG", "Cornell Lossless JPEG License",
		"CPAL-1.0", "Common Public Attribution License 1.0",
		"CPL-1.0", "Common Public License 1.0",
		"CPOL-1.02", "Code Project Open License 1.02",
		"Cronyx", "Cronyx License",
		"Crossword", "Crossword License",
		"CrystalStacker", "CrystalStacker License",
		"CUA-OPL-1.0", "CUA Office Public License v1.0",
		"Cube", "Cube License",
		"C-UDA-1.0", "Computational Use of Data Agreement v1.0",
		"curl", "curl License",
		"DEC-3-Clause", "DEC 3-Clause License",
		"D-FSL-1.0", "Deutsche Freie Software Lizenz",
		"diffmark", "diffmark license",
		"DL-DE-BY-2.0", "Data licence Germany – attribution – version 2.0",
		"DL-DE-ZERO-2.0", "Data licence Germany – zero – version 2.0",
		"DOC", "DOC License",
		"Dotseqn", "Dotseqn License",
		"DRL-1.0", "Detection Rule License 1.0",
		"DRL-1.1", "Detection Rule License 1.1",
		"DSDP", "DSDP License",
		"dtoa", "David M. Gay dtoa License",
		"dvipdfm", "dvipdfm License",
		"ECL-1.0", "Educational Community License v1.0",
		"ECL-2.0", "Educational Community License v2.0",
		"eCos-2.0", "eCos license version 2.0",
		"EFL-1.0", "Eiffel Forum License v1.0",
		"EFL-2.0", "Eiffel Forum License v2.0",
		"eGenix", "eGenix.com Public License 1.1.0",
		"Elastic-2.0", "Elastic License 2.0",
		"Entessa", "Entessa Public License v1.0",
		"EPICS", "EPICS Open License",
		"EPL-1.0", "Eclipse Public License 1.0",
		"EPL-2.0", "Eclipse Public License 2.0",
		"ErlPL-1.1", "Erlang Public License v1.1",
		"etalab-2.0", "Etalab Open License 2.0",
		"EUDatagrid", "EU DataGrid Software License",
		"EUPL-1.0", "European Union Public License 1.0",
		"EUPL-1.1", "European Union Public License 1.1",
		"EUPL-1.2", "European Union Public License 1.2",
		"Eurosym", "Eurosym License",
		"Fair", "Fair License",
		"FBM", "Fuzzy Bitmap License",
		"FDK-AAC", "Fraunhofer FDK AAC Codec Library",
		"Ferguson-Twofish", "Ferguson Twofish License",
		"Frameworx-1.0", "Frameworx Open License 1.0",
		"FreeBSD-DOC", "FreeBSD Documentation License",
		"FreeImage", "FreeImage Public License v1.0",
		"FSFAP", "FSF All Permissive License",
		"FSFAP-no-warranty-disclaimer", "FSF All Permissive License (without Warranty)",
		"FSFUL", "FSF Unlimited License",
		"FSFULLR", "FSF Unlimited License (with License Retention)",
		"FSFULLRWD", "FSF Unlimited License (With License Retention and Warranty Disclaimer)",
		"FTL", "Freetype Project License",
		"Furuseth", "Furuseth License",
		"fwlw", "fwlw License",
		"GCR-docs", "Gnome GCR Documentation License",
		"GD", "GD License",
		"GFDL-1.1", "GNU Free Documentation License v1.1",
		"GFDL-1.1-invariants-only", "GNU Free Documentation License v1.1 only - invariants",
		"GFDL-1.1-invariants-or-later", "GNU Free Documentation License v1.1 or later - invariants",
		"GFDL-1.1-no-invariants-only", "GNU Free Documentation License v1.1 only - no invariants",
		"GFDL-1.1-no-invariants-or-later", "GNU Free Documentation License v1.1 or later - no invariants",
		"GFDL-1.1-only", "GNU Free Documentation License v1.1 only",
		"GFDL-1.1-or-later", "GNU Free Documentation License v1.1 or later",
		"GFDL-1.2", "GNU Free Documentation License v1.2",
		"GFDL-1.2-invariants-only", "GNU Free Documentation License v1.2 only - invariants",
		"GFDL-1.2-invariants-or-later", "GNU Free Documentation License v1.2 or later - invariants",
		"GFDL-1.2-no-invariants-only", "GNU Free Documentation License v1.2 only - no invariants",
		"GFDL-1.2-no-invariants-or-later", "GNU Free Documentation License v1.2 or later - no invariants",
		"GFDL-1.2-only", "GNU Free Documentation License v1.2 only",
		"GFDL-1.2-or-later", "GNU Free Documentation License v1.2 or later",
		"GFDL-1.3", "GNU Free Documentation License v1.3",
		"GFDL-1.3-invariants-only", "GNU Free Documentation License v1.3 only - invariants",
		"GFDL-1.3-invariants-or-later", "GNU Free Documentation License v1.3 or later - invariants",
		"GFDL-1.3-no-invariants-only", "GNU Free Documentation License v1.3 only - no invariants",
		"GFDL-1.3-no-invariants-or-later", "GNU Free Documentation License v1.3 or later - no invariants",
		"GFDL-1.3-only", "GNU Free Documentation License v1.3 only",
		"GFDL-1.3-or-later", "GNU Free Documentation License v1.3 or later",
		"Giftware", "Giftware License",
		"GL2PS", "GL2PS License",
		"Glide", "3dfx Glide License",
		"Glulxe", "Glulxe License",
		"GLWTPL", "Good Luck With That Public License",
		"gnuplot", "gnuplot License",
		"GPL-1.0", "GNU General Public License v1.0 only",
		"GPL-1.0+", "GNU General Public License v1.0 or later",
		"GPL-1.0-only", "GNU General Public License v1.0 only",
		"GPL-1.0-or-later", "GNU General Public License v1.0 or later",
		"GPL-2.0", "GNU General Public License v2.0 only",
		"GPL-2.0+", "GNU General Public License v2.0 or later",
		"GPL-2.0-only", "GNU General Public License v2.0 only",
		"GPL-2.0-or-later", "GNU General Public License v2.0 or later",
		"GPL-2.0-with-autoconf-exception", "GNU General Public License v2.0 w/Autoconf exception",
		"GPL-2.0-with-bison-exception", "GNU General Public License v2.0 w/Bison exception",
		"GPL-2.0-with-classpath-exception", "GNU General Public License v2.0 w/Classpath exception",
		"GPL-2.0-with-font-exception", "GNU General Public License v2.0 w/Font exception",
		"GPL-2.0-with-GCC-exception", "GNU General Public License v2.0 w/GCC Runtime Library exception",
		"GPL-3.0", "GNU General Public License v3.0 only",
		"GPL-3.0+", "GNU General Public License v3.0 or later",
		"GPL-3.0-only", "GNU General Public License v3.0 only",
		"GPL-3.0-or-later", "GNU General Public License v3.0 or later",
		"GPL-3.0-with-autoconf-exception", "GNU General Public License v3.0 w/Autoconf exception",
		"GPL-3.0-with-GCC-exception", "GNU General Public License v3.0 w/GCC Runtime Library exception",
		"Graphics-Gems", "Graphics Gems License",
		"gSOAP-1.3b", "gSOAP Public License v1.3b",
		"gtkbook", "gtkbook License",
		"HaskellReport", "Haskell Language Report License",
		"hdparm", "hdparm License",
		"Hippocratic-2.1", "Hippocratic License 2.1",
		"HP-1986", "Hewlett-Packard 1986 License",
		"HP-1989", "Hewlett-Packard 1989 License",
		"HPND-DEC", "Historical Permission Notice and Disclaimer - DEC variant",
		"HPND-doc", "Historical Permission Notice and Disclaimer - documentation variant",
		"HPND-doc-sell", "Historical Permission Notice and Disclaimer - documentation sell variant",
		"HPND-export-US", "HPND with US Government export control warning",
		"HPND-export-US-modify", "HPND with US Government export control warning and modification rqmt",
		"HPND-Fenneberg-Livingston", "Historical Permission Notice and Disclaimer - Fenneberg-Livingston variant",
		"HPND", "Historical Permission Notice and Disclaimer",
		"HPND-INRIA-IMAG", "Historical Permission Notice and Disclaimer    - INRIA-IMAG variant",
		"HPND-Kevlin-Henney", "Historical Permission Notice and Disclaimer - Kevlin Henney variant",
		"HPND-Markus-Kuhn", "Historical Permission Notice and Disclaimer - Markus Kuhn variant",
		"HPND-MIT-disclaimer", "Historical Permission Notice and Disclaimer with MIT disclaimer",
		"HPND-Pbmplus", "Historical Permission Notice and Disclaimer - Pbmplus variant",
		"HPND-sell-MIT-disclaimer-xserver", "Historical Permission Notice and Disclaimer - sell xserver variant with MIT disclaimer",
		"HPND-sell-regexpr", "Historical Permission Notice and Disclaimer - sell regexpr variant",
		"HPND-sell-variant", "Historical Permission Notice and Disclaimer - sell variant",
		"HPND-sell-variant-MIT-disclaimer", "HPND sell variant with MIT disclaimer",
		"HPND-UC", "Historical Permission Notice and Disclaimer - University of California variant",
		"HTMLTIDY", "HTML Tidy License",
		"IBM-pibs", "IBM PowerPC Initialization and Boot Software",
		"ICU", "ICU License",
		"Identifier", "Full name",
		"IEC-Code-Components-EULA", "IEC    Code Components End-user licence agreement",
		"IJG", "Independent JPEG Group License",
		"IJG-short", "Independent JPEG Group License - short",
		"ImageMagick", "ImageMagick License",
		"iMatix", "iMatix Standard Function Library Agreement",
		"Imlib2", "Imlib2 License",
		"Info-ZIP", "Info-ZIP License",
		"Inner-Net-2.0", "Inner Net License v2.0",
		"Intel-ACPI", "Intel ACPI Software License Agreement",
		"Intel", "Intel Open Source License",
		"Interbase-1.0", "Interbase Public License v1.0",
		"IPA", "IPA Font License",
		"IPL-1.0", "IBM Public License v1.0",
		"ISC", "ISC License",
		"ISC-Veillard", "ISC Veillard variant",
		"Jam", "Jam License",
		"JasPer-2.0", "JasPer License",
		"JPL-image", "JPL Image Use Policy",
		"JPNIC", "Japan Network Information Center License",
		"JSON", "JSON License",
		"Kastrup", "Kastrup License",
		"Kazlib", "Kazlib License",
		"Knuth-CTAN", "Knuth CTAN License",
		"LAL-1.2", "Licence Art Libre 1.2",
		"LAL-1.3", "Licence Art Libre 1.3",
		"Latex2e", "Latex2e License",
		"Latex2e-translated-notice", "Latex2e with translated notice permission",
		"Leptonica", "Leptonica License",
		"LGPL-2.0", "GNU Library General Public License v2 only",
		"LGPL-2.0+", "GNU Library General Public License v2 or later",
		"LGPL-2.0-only", "GNU Library General Public License v2 only",
		"LGPL-2.0-or-later", "GNU Library General Public License v2 or later",
		"LGPL-2.1", "GNU Lesser General Public License v2.1 only",
		"LGPL-2.1+", "GNU Lesser General Public License v2.1 or later",
		"LGPL-2.1-only", "GNU Lesser General Public License v2.1 only",
		"LGPL-2.1-or-later", "GNU Lesser General Public License v2.1 or later",
		"LGPL-3.0", "GNU Lesser General Public License v3.0 only",
		"LGPL-3.0+", "GNU Lesser General Public License v3.0 or later",
		"LGPL-3.0-only", "GNU Lesser General Public License v3.0 only",
		"LGPL-3.0-or-later", "GNU Lesser General Public License v3.0 or later",
		"LGPLLR", "Lesser General Public License For Linguistic Resources",
		"libpng-2.0", "PNG Reference Library version 2",
		"Libpng", "libpng License",
		"libselinux-1.0", "libselinux public domain notice",
		"libtiff", "libtiff License",
		"libutil-David-Nugent", "libutil David Nugent License",
		"LiLiQ-P-1.1", "Licence Libre du Québec – Permissive version 1.1",
		"LiLiQ-R-1.1", "Licence Libre du Québec – Réciprocité version 1.1",
		"LiLiQ-Rplus-1.1", "Licence Libre du Québec – Réciprocité forte version 1.1",
		"Linux-man-pages-1-para", "Linux man-pages - 1 paragraph",
		"Linux-man-pages-copyleft-2-para", "Linux man-pages Copyleft - 2 paragraphs",
		"Linux-man-pages-copyleft", "Linux man-pages Copyleft",
		"Linux-man-pages-copyleft-var", "Linux man-pages Copyleft Variant",
		"Linux-OpenIB", "Linux Kernel Variant of OpenIB.org license",
		"LOOP", "Common Lisp LOOP License",
		"LPD-document", "LPD Documentation License",
		"LPL-1.02", "Lucent Public License v1.02",
		"LPL-1.0", "Lucent Public License Version 1.0",
		"LPPL-1.0", "LaTeX Project Public License v1.0",
		"LPPL-1.1", "LaTeX Project Public License v1.1",
		"LPPL-1.2", "LaTeX Project Public License v1.2",
		"LPPL-1.3a", "LaTeX Project Public License v1.3a",
		"LPPL-1.3c", "LaTeX Project Public License v1.3c",
		"lsof", "lsof License",
		"Lucida-Bitmap-Fonts", "Lucida Bitmap Fonts License",
		"LZMA-SDK-9.11-to-9.20", "LZMA SDK License (versions 9.11 to 9.20)",
		"LZMA-SDK-9.22", "LZMA SDK License (versions 9.22 and beyond)",
		"Mackerras-3-Clause-acknowledgment", "Mackerras 3-Clause - acknowledgment variant",
		"Mackerras-3-Clause", "Mackerras 3-Clause License",
		"magaz", "magaz License",
		"mailprio", "mailprio License",
		"MakeIndex", "MakeIndex License",
		"Martin-Birgmeier", "Martin Birgmeier License",
		"McPhee-slideshow", "McPhee Slideshow License",
		"metamail", "metamail License",
		"Minpack", "Minpack License",
		"MirOS", "The MirOS Licence",
		"MIT-0", "MIT No Attribution",
		"MIT-advertising", "Enlightenment License (e16)",
		"MIT-CMU", "CMU License",
		"MIT-enna", "enna License",
		"MIT-feh", "feh License",
		"MIT-Festival", "MIT Festival Variant",
		"MIT", "MIT License",
		"MIT-Modern-Variant", "MIT License Modern Variant",
		"MITNFA", "MIT +no-false-attribs license",
		"MIT-open-group", "MIT Open Group variant",
		"MIT-testregex", "MIT testregex Variant",
		"MIT-Wu", "MIT Tom Wu Variant",
		"MMIXware", "MMIXware License",
		"Motosoto", "Motosoto License",
		"MPEG-SSG", "MPEG Software Simulation",
		"mpich2", "mpich2 License",
		"mpi-permissive", "mpi Permissive License",
		"MPL-1.0", "Mozilla Public License 1.0",
		"MPL-1.1", "Mozilla Public License 1.1",
		"MPL-2.0", "Mozilla Public License 2.0",
		"MPL-2.0-no-copyleft-exception", "Mozilla Public License 2.0 (no copyleft exception)",
		"mplus", "mplus Font License",
		"MS-LPL", "Microsoft Limited Public License",
		"MS-PL", "Microsoft Public License",
		"MS-RL", "Microsoft Reciprocal License",
		"MTLL", "Matrix Template Library License",
		"MulanPSL-1.0", "Mulan Permissive Software License, Version 1",
		"MulanPSL-2.0", "Mulan Permissive Software License, Version 2",
		"Multics", "Multics License",
		"Mup", "Mup License",
		"NAIST-2003", "Nara Institute of Science and Technology License (2003)",
		"NASA-1.3", "NASA Open Source Agreement 1.3",
		"Naumen", "Naumen Public License",
		"NBPL-1.0", "Net Boolean Public License v1",
		"NCGL-UK-2.0", "Non-Commercial Government Licence",
		"NCSA", "University of Illinois/NCSA Open Source License",
		"NetCDF", "NetCDF license",
		"Net-SNMP", "Net-SNMP License",
		"Newsletr", "Newsletr License",
		"NGPL", "Nethack General Public License",
		"NICTA-1.0", "NICTA Public Software License, Version 1.0",
		"NIST-PD-fallback", "NIST Public Domain Notice with license fallback",
		"NIST-PD", "NIST Public Domain Notice",
		"NIST-Software", "NIST Software License",
		"NLOD-1.0", "Norwegian Licence for Open Government Data (NLOD) 1.0",
		"NLOD-2.0", "Norwegian Licence for Open Government Data (NLOD) 2.0",
		"NLPL", "No Limit Public License",
		"Nokia", "Nokia Open Source License",
		"NOSL", "Netizen Open Source License",
		"Noweb", "Noweb License",
		"NPL-1.0", "Netscape Public License v1.0",
		"NPL-1.1", "Netscape Public License v1.1",
		"NPOSL-3.0", "Non-Profit Open Software License 3.0",
		"NRL", "NRL License",
		"NTP-0", "NTP No Attribution",
		"NTP", "NTP License",
		"Nunit", "Nunit License",
		"OCCT-PL", "Open CASCADE Technology Public License",
		"OCLC-2.0", "OCLC Research Public License 2.0",
		"ODbL-1.0", "Open Data Commons Open Database License v1.0",
		"ODC-By-1.0", "Open Data Commons Attribution License v1.0",
		"OFFIS", "OFFIS License",
		"OFL-1.0-no-RFN", "SIL Open Font License 1.0 with no Reserved Font Name",
		"OFL-1.0-RFN", "SIL Open Font License 1.0 with Reserved Font Name",
		"OFL-1.0", "SIL Open Font License 1.0",
		"OFL-1.1-no-RFN", "SIL Open Font License 1.1 with no Reserved Font Name",
		"OFL-1.1-RFN", "SIL Open Font License 1.1 with Reserved Font Name",
		"OFL-1.1", "SIL Open Font License 1.1",
		"OGC-1.0", "OGC Software License, Version 1.0",
		"OGDL-Taiwan-1.0", "Taiwan Open Government Data License, version 1.0",
		"OGL-Canada-2.0", "Open Government Licence - Canada",
		"OGL-UK-1.0", "Open Government Licence v1.0",
		"OGL-UK-2.0", "Open Government Licence v2.0",
		"OGL-UK-3.0", "Open Government Licence v3.0",
		"OGTSL", "Open Group Test Suite License",
		"OLDAP-1.1", "Open LDAP Public License v1.1",
		"OLDAP-1.2", "Open LDAP Public License v1.2",
		"OLDAP-1.3", "Open LDAP Public License v1.3",
		"OLDAP-1.4", "Open LDAP Public License v1.4",
		"OLDAP-2.0.1", "Open LDAP Public License v2.0.1",
		"OLDAP-2.0", "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)",
		"OLDAP-2.1", "Open LDAP Public License v2.1",
		"OLDAP-2.2.1", "Open LDAP Public License v2.2.1",
		"OLDAP-2.2.2", "Open LDAP Public License 2.2.2",
		"OLDAP-2.2", "Open LDAP Public License v2.2",
		"OLDAP-2.3", "Open LDAP Public License v2.3",
		"OLDAP-2.4", "Open LDAP Public License v2.4",
		"OLDAP-2.5", "Open LDAP Public License v2.5",
		"OLDAP-2.6", "Open LDAP Public License v2.6",
		"OLDAP-2.7", "Open LDAP Public License v2.7",
		"OLDAP-2.8", "Open LDAP Public License v2.8",
		"OLFL-1.3", "Open Logistics Foundation License Version 1.3",
		"OML", "Open Market License",
		"OpenPBS-2.3", "OpenPBS v2.3 Software License",
		"OpenSSL", "OpenSSL License",
		"OpenSSL-standalone", "OpenSSL License - standalone",
		"OpenVision", "OpenVision License",
		"OPL-1.0", "Open Public License v1.0",
		"OPL-UK-3.0", "United    Kingdom Open Parliament Licence v3.0",
		"OPUBL-1.0", "Open Publication License v1.0",
		"OSET-PL-2.1", "OSET Public License version 2.1",
		"OSL-1.0", "Open Software License 1.0",
		"OSL-1.1", "Open Software License 1.1",
		"OSL-2.0", "Open Software License 2.0",
		"OSL-2.1", "Open Software License 2.1",
		"OSL-3.0", "Open Software License 3.0",
		"O-UDA-1.0", "Open Use of Data Agreement v1.0",
		"PADL", "PADL License",
		"Parity-6.0.0", "The Parity Public License 6.0.0",
		"Parity-7.0.0", "The Parity Public License 7.0.0",
		"PDDL-1.0", "Open Data Commons Public Domain Dedication & License 1.0",
		"PHP-3.01", "PHP License v3.01",
		"PHP-3.0", "PHP License v3.0",
		"Pixar", "Pixar License",
		"Plexus", "Plexus Classworlds License",
		"pnmstitch", "pnmstitch License",
		"PolyForm-Noncommercial-1.0.0", "PolyForm Noncommercial License 1.0.0",
		"PolyForm-Small-Business-1.0.0", "PolyForm Small Business License 1.0.0",
		"PostgreSQL", "PostgreSQL License",
		"PSF-2.0", "Python Software Foundation License 2.0",
		"psfrag", "psfrag License",
		"psutils", "psutils License",
		"Python-2.0.1", "Python License 2.0.1",
		"Python-2.0", "Python License 2.0",
		"python-ldap", "Python ldap License",
		"Qhull", "Qhull License",
		"QPL-1.0-INRIA-2004", "Q Public License 1.0 - INRIA 2004 variant",
		"QPL-1.0", "Q Public License 1.0",
		"radvd", "radvd License",
		"Rdisc", "Rdisc License",
		"RHeCos-1.1", "Red Hat eCos Public License v1.1",
		"RPL-1.1", "Reciprocal Public License 1.1",
		"RPL-1.5", "Reciprocal Public License 1.5",
		"RPSL-1.0", "RealNetworks Public Source License v1.0",
		"RSA-MD", "RSA Message-Digest License",
		"RSCPL", "Ricoh Source Code Public License",
		"Ruby", "Ruby License",
		"Saxpath", "Saxpath License",
		"SAX-PD-2.0", "Sax Public Domain Notice 2.0",
		"SAX-PD", "Sax Public Domain Notice",
		"SCEA", "SCEA Shared Source License",
		"SchemeReport", "Scheme Language Report License",
		"Sendmail-8.23", "Sendmail License 8.23",
		"Sendmail", "Sendmail License",
		"SGI-B-1.0", "SGI Free Software License B v1.0",
		"SGI-B-1.1", "SGI Free Software License B v1.1",
		"SGI-B-2.0", "SGI Free Software License B v2.0",
		"SGI-OpenGL", "SGI OpenGL License",
		"SGP4", "SGP4 Permission Notice",
		"SHL-0.51", "Solderpad Hardware License, Version 0.51",
		"SHL-0.5", "Solderpad Hardware License v0.5",
		"SimPL-2.0", "Simple Public License 2.0",
		"SISSL-1.2", "Sun Industry Standards Source License v1.2",
		"SISSL", "Sun Industry Standards Source License v1.1",
		"Sleepycat", "Sleepycat License",
		"SL", "SL License",
		"SMLNJ", "Standard ML of New Jersey License",
		"SMPPL", "Secure Messaging Protocol Public License",
		"SNIA", "SNIA Public License 1.1",
		"snprintf", "snprintf License",
		"softSurfer", "softSurfer License",
		"Soundex", "Soundex License",
		"Spencer-86", "Spencer License 86",
		"Spencer-94", "Spencer License 94",
		"Spencer-99", "Spencer License 99",
		"SPL-1.0", "Sun Public License v1.0",
		"ssh-keyscan", "ssh-keyscan License",
		"SSH-OpenSSH", "SSH OpenSSH license",
		"SSH-short", "SSH short notice",
		"SSLeay-standalone", "SSLeay License - standalone",
		"SSPL-1.0", "Server Side Public License, v 1",
		"StandardML-NJ", "Standard ML of New Jersey License",
		"SugarCRM-1.1.3", "SugarCRM Public License v1.1.3",
		"Sun-PPP", "Sun PPP License",
		"SunPro", "SunPro License",
		"SWL", "Scheme Widget Library (SWL) Software License Agreement",
		"swrule", "swrule License",
		"Symlinks", "Symlinks License",
		"TAPR-OHL-1.0", "TAPR Open Hardware License v1.0",
		"TCL", "TCL/TK License",
		"TCP-wrappers", "TCP Wrappers License",
		"TermReadKey", "TermReadKey License",
		"TGPPL-1.0", "Transitive Grace Period Public Licence 1.0",
		"TMate", "TMate Open Source License",
		"TORQUE-1.1", "TORQUE v2.5+ Software License v1.1",
		"TOSL", "Trusster Open Source License",
		"TPDL", "Time::ParseDate License",
		"TPL-1.0", "THOR Public License 1.0",
		"TTWL", "Text-Tabs+Wrap License",
		"TTYP0", "TTYP0 License",
		"TU-Berlin-1.0", "Technische Universitaet Berlin License 1.0",
		"TU-Berlin-2.0", "Technische Universitaet Berlin License 2.0",
		"UCAR", "UCAR License",
		"UCL-1.0", "Upstream Compatibility License v1.0",
		"ulem", "ulem License",
		"UMich-Merit", "Michigan/Merit Networks License",
		"Unicode-3.0", "Unicode License v3",
		"Unicode-DFS-2015", "Unicode License Agreement - Data Files and Software (2015)",
		"Unicode-DFS-2016", "Unicode License Agreement - Data Files and Software (2016)",
		"Unicode-TOU", "Unicode Terms of Use",
		"UnixCrypt", "UnixCrypt License",
		"Unlicense", "The Unlicense",
		"UPL-1.0", "Universal Permissive License v1.0",
		"URT-RLE", "Utah Raster Toolkit Run Length Encoded License",
		"Vim", "Vim License",
		"VOSTROM", "VOSTROM Public License for Open Source",
		"VSL-1.0", "Vovida Software License v1.0",
		"W3C-19980720", "W3C Software Notice and License (1998-07-20)",
		"W3C-20150513", "W3C Software Notice and Document License (2015-05-13)",
		"W3C", "W3C Software Notice and License (2002-12-31)",
		"w3m", "w3m License",
		"Watcom-1.0", "Sybase Open Watcom Public License 1.0",
		"Widget-Workshop", "Widget Workshop License",
		"Wsuipa", "Wsuipa License",
		"WTFPL", "Do What The F*ck You Want To Public License",
		"wxWindows", "wxWindows Library License",
		"X11-distribute-modifications-variant", "X11 License Distribution Modification Variant",
		"X11", "X11 License",
		"Xdebug-1.03", "Xdebug License v 1.03",
		"Xerox", "Xerox License",
		"Xfig", "Xfig License",
		"XFree86-1.1", "XFree86 License 1.1",
		"xinetd", "xinetd License",
		"xkeyboard-config-Zinoviev", "xkeyboard-config Zinoviev License",
		"xlock", "xlock License",
		"Xnet", "X.Net License",
		"xpp", "XPP License",
		"XSkat", "XSkat License",
		"YPL-1.0", "Yahoo! Public License v1.0",
		"YPL-1.1", "Yahoo! Public License v1.1",
		"Zed", "Zed License",
		"Zeeff", "Zeeff License",
		"Zend-2.0", "Zend License v2.0",
		"Zimbra-1.3", "Zimbra Public License v1.3",
		"Zimbra-1.4", "Zimbra Public License v1.4",
		"zlib-acknowledgement", "zlib/libpng License with Acknowledgement",
		"Zlib", "zlib License",
		"ZPL-1.1", "Zope Public License 1.1",
		"ZPL-2.0", "Zope Public License 2.0",
		"ZPL-2.1", "Zope Public License 2.1",
	).Tag("licenses")
}
