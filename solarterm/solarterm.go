package solarterm

import (
	"time"

	"github.com/codus-me/chinese-calendar-golang/utils"
)

// Solarterm 节气
type Solarterm struct {
	index int64
}

// SolartermFromYear 支持的最早年份
const SolartermFromYear = 1904

// SolartermToYear 支持的最晚年份
const SolartermToYear = 2024

var solartermTimestamp = [...]int64{
	// 1904
	-2082349399, -2081077349, -2079801372, -2078519728, -2077229320, -2075929307,
	-2074617687, -2073295088, -2071960898, -2070617473, -2069265548, -2067908921,
	-2066549299, -2065191025, -2063836089, -2062488216, -2061148922, -2059820386,
	-2058503062, -2057197251, -2055902092, -2054616234, -2053337663, -2052063944,
	// 1905
	-2050792350, -2049520060, -2048244225, -2046962316, -2045672037, -2044371718,
	-2043060299, -2041737341, -2040403522, -2039059690, -2037708350, -2036351281,
	-2034992362, -2033633626, -2032279346, -2030931046, -2029592257, -2028263366,
	-2026946390, -2025640293, -2024345384, -2023059279, -2021780923, -2020506948,
	// 1906
	-2019235560, -2017962973, -2016687336, -2015405104, -2014115008, -2012814408,
	-2011503140, -2010180028, -2008846271, -2007502487, -2006151050, -2004794277,
	-2003435069, -2002076833, -2000722091, -1999374377, -1998035017, -1996706693,
	-1995389099, -1994083511, -1992787983, -1991502373, -1990223443, -1988950016,
	// 1907
	-1987678126, -1986406162, -1985130078, -1983848511, -1982557986, -1981258034,
	-1979946329, -1978623786, -1977289605, -1975946230, -1974594448, -1973237845,
	-1971878472, -1970520141, -1969165460, -1967817414, -1966478298, -1965149489,
	-1963832262, -1962526133, -1961231046, -1959944901, -1958666457, -1957392533,
	// 1908
	-1956121156, -1954848737, -1953573183, -1952291177, -1951001192, -1949700770,
	-1948389615, -1947066527, -1945732903, -1944388917, -1943037660, -1941680459,
	-1940321521, -1938962753, -1937608391, -1936260171, -1934921250, -1933592494,
	-1932275337, -1930969381, -1929674266, -1928388307, -1927109765, -1925835975,
	// 1909
	-1924564462, -1923292119, -1922016419, -1920734468, -1919444316, -1918143984,
	-1916832592, -1915509691, -1914175707, -1912832064, -1911480321, -1910123626,
	-1908764120, -1907405932, -1906051012, -1904703349, -1903363964, -1902035687,
	-1900718167, -1899412605, -1898117175, -1896831547, -1895552676, -1894279179,
	// 1910
	-1893007289, -1891735227, -1890459124, -1889177482, -1887886980, -1886586999,
	-1885275395, -1883952830, -1882618816, -1881275370, -1879923801, -1878567061,
	-1877207921, -1875849407, -1874494956, -1873146749, -1871807859, -1870478948,
	-1869161924, -1867855723, -1866560787, -1865274543, -1863996184, -1862722097,
	// 1911
	-1861450750, -1860178121, -1858902590, -1857620392, -1856330482, -1855029953,
	-1853718945, -1852395859, -1851062399, -1849718503, -1848367347, -1847010292,
	-1845651329, -1844292708, -1842938159, -1841590046, -1840250828, -1838922175,
	-1837604731, -1836298933, -1835003602, -1833717866, -1832439165, -1831165628,
	// 1912
	-1829893969, -1828621870, -1827346005, -1826064281, -1824773955, -1823473853,
	-1822162314, -1820839665, -1819505578, -1818162172, -1816810346, -1815453784,
	-1814094193, -1812735977, -1811380963, -1810033115, -1808693651, -1807365109,
	-1806047583, -1804741785, -1803446466, -1802160695, -1800882049, -1799608499,
	// 1913
	-1798336903, -1797064830, -1795789017, -1794507322, -1793217035, -1791916893,
	-1790605415, -1789282592, -1787948686, -1786604973, -1785253558, -1783896595,
	-1782537628, -1781178940, -1779824616, -1778476272, -1777137422, -1775808399,
	-1774491344, -1773185072, -1771890102, -1770603853, -1769325506, -1768051479,
	// 1914
	-1766780199, -1765507664, -1764232219, -1762950103, -1761660229, -1760359736,
	-1759048667, -1757725585, -1756391974, -1755048124, -1753696786, -1752339885,
	-1750980753, -1749622372, -1748267675, -1746919814, -1745580445, -1744251969,
	-1742934309, -1741628563, -1740332937, -1739047183, -1737768182, -1736494668,
	// 1915
	-1735222795, -1733950842, -1732674888, -1731393434, -1730103124, -1728803347,
	-1727491867, -1726169495, -1724835456, -1723492198, -1722140416, -1720783863,
	-1719424359, -1718066042, -1716711165, -1715363127, -1714023802, -1712694999,
	-1711377571, -1710071444, -1708776164, -1707490016, -1706211390, -1704937479,
	// 1916
	-1703665955, -1702393607, -1701117979, -1699836136, -1698546171, -1697246000,
	-1695934940, -1694612135, -1693278622, -1691934857, -1690583662, -1689226537,
	-1687867584, -1686508726, -1685154297, -1683805881, -1682466891, -1681137910,
	-1679820720, -1678514558, -1677219450, -1675933317, -1674654810, -1673380868,
	// 1917
	-1672109410, -1670836938, -1669561323, -1668279289, -1666989283, -1665688938,
	-1664377774, -1663054925, -1661721225, -1660377654, -1659026174, -1657669507,
	-1656310150, -1654951896, -1653596954, -1652249144, -1650909602, -1649581156,
	-1648263437, -1646957749, -1645662157, -1644376481, -1643097514, -1641824035,
	// 1918
	-1640552106, -1639280099, -1638003987, -1636722417, -1635431923, -1634132043,
	-1632820464, -1631498054, -1630164086, -1628820851, -1627469322, -1626112806,
	-1624753655, -1623395302, -1622040746, -1620692568, -1619353470, -1618024462,
	-1616707181, -1615400835, -1614105671, -1612819325, -1611540821, -1610266725,
	// 1919
	-1608995325, -1607722773, -1606447251, -1605165171, -1603875290, -1602574876,
	-1601263899, -1599940905, -1598607502, -1597263676, -1595912626, -1594555614,
	-1593196795, -1591838161, -1590483746, -1589135532, -1587796373, -1586467513,
	-1585150031, -1583843951, -1582548531, -1581262509, -1579983749, -1578709996,
	// 1920
	-1577438369, -1576166157, -1574890428, -1573608677, -1572318550, -1571018456,
	-1569707114, -1568384458, -1567050524, -1565707096, -1564355375, -1562998812,
	-1561639282, -1560281105, -1558926103, -1557578320, -1556238801, -1554910307,
	-1553592640, -1552286830, -1550991291, -1549705459, -1548426563, -1547152959,
	// 1921
	-1545881154, -1544609095, -1543333165, -1542051579, -1540761267, -1539461311,
	-1538149846, -1536827230, -1535493306, -1534149758, -1532798271, -1531441419,
	-1530082361, -1528723744, -1527369356, -1526021056, -1524682186, -1523353181,
	-1522036129, -1520729829, -1519434837, -1518148500, -1516870084, -1515595926,
	// 1922
	-1514324554, -1513051895, -1511776386, -1510494203, -1509204345, -1507903861,
	-1506592896, -1505269865, -1503936407, -1502592568, -1501241365, -1499884383,
	-1498525335, -1497166804, -1495812157, -1494464141, -1493124818, -1491796229,
	-1490478636, -1489172830, -1487877290, -1486591496, -1485312571, -1484038998,
	// 1923
	-1482767170, -1481495122, -1480219191, -1478937628, -1477647346, -1476347492,
	-1475036069, -1473713686, -1472379727, -1471036511, -1469684770, -1468328267,
	-1466968698, -1465610398, -1464255357, -1462907319, -1461567796, -1460239015,
	-1458921424, -1457615375, -1456320004, -1455034006, -1453755346, -1452481624,
	// 1924
	-1451210080, -1449937910, -1448662234, -1447380528, -1446090470, -1444790397,
	-1443479215, -1442156491, -1440822865, -1439479180, -1438127910, -1436770842,
	-1435411836, -1434052950, -1432698460, -1431349915, -1430010855, -1428681691,
	-1427364453, -1426058121, -1424763032, -1423476803, -1422198405, -1420924459,
	// 1925
	-1419653187, -1418380771, -1417105371, -1415823394, -1414533581, -1413233242,
	-1411922219, -1410599301, -1409265695, -1407921991, -1406570581, -1405213766,
	-1403854464, -1402496069, -1401141130, -1399793172, -1398453556, -1397124959,
	-1395807117, -1394501298, -1393205591, -1391919843, -1390640832, -1389367375,
	// 1926
	-1388095512, -1386823627, -1385547678, -1384266295, -1382975997, -1381676310,
	-1380364878, -1379042613, -1377708677, -1376365514, -1375013885, -1373657387,
	-1372298051, -1370939704, -1369584939, -1368236759, -1366897446, -1365568408,
	-1364250911, -1362944511, -1361649145, -1360362757, -1359084091, -1357810013,
	// 1927
	-1356538534, -1355266103, -1353990611, -1352708761, -1351419000, -1350118876,
	-1348808052, -1347485319, -1346152034, -1344808348, -1343457334, -1342100293,
	-1340741427, -1339382624, -1338028140, -1336679701, -1335340499, -1334011412,
	-1332693915, -1331387618, -1330092200, -1328805979, -1327527237, -1326253311,
	// 1928
	-1324981744, -1323709418, -1322433828, -1321152057, -1319862173, -1318562156,
	-1317251133, -1315928603, -1314594993, -1313251661, -1311900169, -1310543614,
	-1309184138, -1307825857, -1306470740, -1305122802, -1303783082, -1302454459,
	-1301136591, -1299830712, -1298535007, -1297249163, -1295970137, -1294696557,
	// 1929
	-1293424653, -1292152642, -1290876652, -1289595163, -1288304852, -1287005078,
	-1285693688, -1284371345, -1283037540, -1281694304, -1280342908, -1278986321,
	-1277627258, -1276268763, -1274914236, -1273565884, -1272226781, -1270897620,
	-1269580335, -1268273876, -1266978716, -1265692283, -1264413783, -1263139607,
	// 1930
	-1261868217, -1260595594, -1259320105, -1258037989, -1256748182, -1255447794,
	-1254136934, -1252814032, -1251480759, -1250137062, -1248786101, -1247429219,
	-1246070406, -1244711874, -1243357373, -1242009218, -1240669895, -1239341044,
	-1238023348, -1236717238, -1235421586, -1234135539, -1232856574, -1231582844,
	// 1931
	-1230311078, -1229038963, -1227763174, -1226481598, -1225191490, -1223891643,
	-1222580394, -1221258036, -1219924248, -1218581103, -1217229522, -1215873146,
	-1214513691, -1213155543, -1211800528, -1210452609, -1209112986, -1207784228,
	-1206466411, -1205160293, -1203864627, -1202578537, -1201299600, -1200025843,
	// 1932
	-1198754111, -1197482008, -1196206253, -1194924711, -1193634650, -1192334795,
	-1191023628, -1189701124, -1188367495, -1187024009, -1185672734, -1184315844,
	-1182956862, -1181598117, -1180243685, -1178895219, -1177556219, -1176227039,
	-1174909813, -1173603357, -1172308206, -1171021778, -1169743278, -1168469125,
	// 1933
	-1167197774, -1165925214, -1164649817, -1163367797, -1162078086, -1160777785,
	-1159466936, -1158144070, -1156810656, -1155466955, -1154115717, -1152758849,
	-1151399696, -1150041233, -1148686427, -1147338423, -1145998917, -1144670301,
	-1143352533, -1142046686, -1140750992, -1139465169, -1138186108, -1136912525,
	// 1934
	-1135640581, -1134368558, -1133092556, -1131811081, -1130520799, -1129221107,
	-1127909760, -1126587574, -1125253740, -1123910694, -1122559106, -1121202716,
	-1119843323, -1118485060, -1117130176, -1115782077, -1114442636, -1113113697,
	-1111796109, -1110489833, -1109194406, -1107908148, -1106629417, -1105355447,
	// 1935
	-1104083870, -1102811514, -1101535892, -1100254099, -1098964208, -1097664157,
	-1096353240, -1095030617, -1093697298, -1092353733, -1091002726, -1089645752,
	-1088286893, -1086928050, -1085573555, -1084224987, -1082885779, -1081556540,
	-1080239085, -1078932675, -1077637370, -1076351099, -1075072530, -1073798596,
	// 1936
	-1072527220, -1071254881, -1069979456, -1068697631, -1067407863, -1066107742,
	-1064796803, -1063474141, -1062140613, -1060797155, -1059445756, -1058089100,
	-1056729696, -1055371324, -1054016204, -1052668164, -1051328356, -1049999633,
	-1048681641, -1047375703, -1046079905, -1044794084, -1043515048, -1042241582,
	// 1937
	-1040969753, -1039697917, -1038422039, -1037140729, -1035850503, -1034550863,
	-1033239479, -1031917214, -1030583327, -1029240134, -1027888591, -1026532042,
	-1025172804, -1023814347, -1022459641, -1021111296, -1019772001, -1018442794,
	-1017125312, -1015818780, -1014523455, -1013236984, -1011958396, -1010684276,
	// 1938
	-1009412906, -1008140452, -1006865079, -1005583206, -1004293555, -1002993401,
	-1001682666, -1000359904, -999026679, -997682983, -996331998, -994974985,
	-993616116, -992257376, -990902837, -989554452, -988215113, -986886033,
	-985568319, -984261978, -982966308, -981680049, -980401098, -979127214,
	// 1939
	-977855542, -976583373, -975307787, -974026262, -972736449, -971436716,
	-970125780, -968803520, -967469965, -966126832, -964775333, -963418869,
	-962059330, -960701033, -959345819, -957997759, -956657904, -955329061,
	-954011027, -952704874, -951409012, -950122916, -948843805, -947570070,
	// 1940
	-946298205, -945026178, -943750366, -942468991, -941178972, -939879386,
	-938568333, -937246152, -935912629, -934569422, -933221759, -931865017,
	-930505916, -929147152, -927792500, -926443878, -925104636, -923775258,
	-922454249, -921147629, -919852379, -918565845, -917287307, -916013096,
	// 1941
	-914741741, -913469160, -912193789, -910911790, -909622164, -908325544,
	-907014867, -905692135, -904358969, -903015393, -901664408, -900307568,
	-898948578, -897589998, -896235215, -894886957, -893547339, -892218405,
	-890896875, -889590740, -888294928, -887008906, -885729815, -884456125,
	// 1942
	-883184233, -881912164, -880636264, -879354775, -878064623, -876764947,
	-875453753, -874131638, -872797977, -871455071, -870103640, -868747421,
	-867388088, -866029954, -864674984, -863326913, -861987240, -860658220,
	-859340302, -858033893, -856738136, -855451783, -854172801, -852898840,
	// 1943
	-851627122, -850354881, -849079210, -847797603, -846507709, -845207868,
	-843896953, -842574536, -841241224, -839897860, -838546892, -837190096,
	-835831303, -834472561, -833118120, -831769529, -830430321, -829100929,
	-827783397, -826476729, -825181296, -823894732, -822616048, -821341874,
	// 1944
	-820070462, -818797990, -817522639, -816240786, -814951185, -813651097,
	-812340368, -811017738, -809684416, -808340950, -806989744, -805633062,
	-804273837, -802915448, -801560466, -800212408, -798872658, -797543892,
	-796225863, -794919828, -793623905, -792337931, -791058721, -789785091,
	// 1945
	-788513107, -787241158, -785965210, -784683881, -783393688, -782094137,
	-780782857, -779460754, -778126968, -776783951, -775432439, -774076043,
	-772716760, -771358445, -770003667, -768655454, -767316082, -765986985,
	-764669420, -763362957, -762067521, -760781066, -759502317, -758228166,
	// 1946
	-756956599, -755684102, -754408546, -753126670, -751836904, -750536825,
	-749226071, -747903457, -746570300, -745226758, -743875874, -742518939,
	-741160149, -739801377, -738446903, -737098422, -735759159, -734429971,
	-733112360, -731805925, -730510379, -729224036, -727945205, -726671220,
	// 1947
	-725399636, -724127328, -722851797, -721570109, -720280346, -718980465,
	-717669615, -716347260, -715013846, -713670685, -712319356, -710962903,
	-709603481, -708245180, -706889977, -705541892, -704201964, -702873114,
	-701554990, -700248877, -698952961, -697666967, -696387850, -695114257,
	// 1948
	-693842406, -692570514, -691294694, -690013410, -688723335, -687423800,
	-686112644, -684790512, -683456867, -682113741, -680762374, -679405759,
	-678046583, -676687942, -675333211, -673984639, -672645291, -671315889,
	-669998372, -668691704, -667396388, -666109850, -664831319, -663557183,
	// 1949
	-662285907, -661013462, -659738202, -658456348, -657166812, -655866687,
	-654556050, -653233330, -651900170, -650556526, -649205553, -647848599,
	-646489665, -645130961, -643776266, -642427874, -641088319, -639759222,
	-638441308, -637134993, -635839185, -634553010, -633273967, -632000201,
	// 1950
	-630728449, -629456400, -628180735, -626899335, -625609460, -624309878,
	-622998918, -621676838, -620343305, -619000359, -617648929, -616292630,
	-614933195, -613575001, -612219887, -610871812, -609531984, -608202993,
	-606884908, -605578526, -604282589, -602996264, -601717115, -600443216,
	// 1951
	-599171392, -597899293, -596623611, -595342243, -594052424, -592752886,
	-591442070, -590119944, -588786670, -587443503, -586092472, -584735740,
	-583376797, -582017992, -580663385, -579314667, -577975345, -576645824,
	-575328246, -574021463, -572726024, -571439355, -570160679, -568886414,
	// 1952
	-567615029, -566342506, -565067235, -563785405, -562495969, -561195983,
	-559885502, -558563006, -557229958, -555886568, -554535577, -553178845,
	-551819715, -550461146, -549106133, -547757823, -546417969, -545088972,
	-543770846, -542464658, -541168691, -539882643, -538603443, -537329785,
	// 1953
	-536057848, -534785891, -533510017, -532228706, -530938620, -529639137,
	-528328007, -527006044, -525672425, -524329596, -522978196, -521621967,
	-520262667, -518904435, -517549489, -516201254, -514861601, -513532413,
	-512214543, -510907998, -509612317, -508325854, -507046957, -505772890,
	// 1954
	-504501256, -503228911, -501953333, -500671638, -499381866, -498081977,
	-496771232, -495448813, -494115698, -492772348, -491421543, -490064751,
	-488706040, -487347296, -485992849, -484644246, -483304931, -481975492,
	-480657768, -479351027, -478055376, -476768765, -475489908, -474215760,
	// 1955
	-472944263, -471671903, -470396554, -469114887, -467825355, -466525511,
	-465214892, -463892550, -462559345, -461216168, -459865026, -458508552,
	-457149278, -455790957, -454435823, -453087692, -451747720, -450418774,
	-449100496, -447794241, -446498108, -445211966, -443932646, -442658961,
	// 1956
	-441386992, -440115109, -438839286, -437558119, -436268128, -434968777,
	-433657724, -432335789, -431002196, -429659241, -428307846, -426951373,
	-425592117, -424233602, -422878778, -421530302, -420190848, -418861481,
	-417543826, -416237119, -414941621, -413654984, -412376247, -411102006,
	// 1957
	-409830546, -408558055, -407282691, -406000890, -404711360, -403411379,
	-402100837, -400778293, -399445263, -398101742, -396750879, -395393933,
	-394035071, -392676273, -391321639, -389973112, -388633632, -387304400,
	-385986564, -384680112, -383384364, -382098028, -380819014, -379545056,
	// 1958
	-378273312, -377001067, -375725429, -374443877, -373154092, -371854437,
	-370543644, -369221569, -367888239, -366545333, -365194064, -363837785,
	-362478394, -361120173, -359764972, -358416853, -357076879, -355747883,
	-354429663, -353123334, -351827300, -350541067, -349261840, -347988035,
	// 1959
	-346716119, -345444088, -344168290, -342886970, -341597029, -340297556,
	-338986642, -337664635, -336331299, -334988297, -333637222, -332280645,
	-330921639, -329562906, -328208185, -326859420, -325519956, -324190323,
	-322872638, -321565763, -320270297, -318983605, -317704982, -316430759,
	// 1960
	-315159467, -313887010, -312611821, -311330032, -310040642, -308740648,
	-307430192, -306107650, -304774646, -303431189, -302080283, -300723462,
	-299364437, -298005750, -296650809, -295302330, -293962467, -292633257,
	-291315063, -290008671, -288712657, -287426476, -286147313, -284873622,
	// 1961
	-283601821, -282329910, -281054231, -279772987, -278483091, -277183643,
	-275872639, -274550664, -273217091, -271874231, -270522804, -269166560,
	-267807169, -266448951, -265093865, -263745654, -262405815, -261076622,
	-259758515, -258451930, -257156001, -255869513, -254590421, -253316410,
	// 1962
	-252044682, -250772507, -249496943, -248215512, -246925818, -245626216,
	-244315535, -242993350, -241660223, -240317006, -238966122, -237609355,
	-236250537, -234891728, -233537186, -232188461, -230849089, -229519497,
	-228201734, -226894810, -225599118, -224312309, -223033423, -221759107,
	// 1963
	-220487632, -219215187, -217939955, -216658307, -215368995, -214069247,
	-212758908, -211436662, -210103712, -208760545, -207409567, -206052991,
	-204693774, -203335277, -201980105, -200631779, -199291720, -197962619,
	-196644251, -195337892, -194041678, -192755456, -191476060, -190202305,
	// 1964
	-188930275, -187658351, -186382514, -185101370, -183811447, -182512216,
	-181201306, -179879580, -178546143, -177203418, -175852098, -174495793,
	-173136471, -171778032, -170423022, -169074527, -167734822, -166405392,
	-165087498, -163780749, -162485077, -161198449, -159919594, -158645406,
	// 1965
	-157373857, -156101443, -154826007, -153544303, -152254733, -150954884,
	-149644365, -148322003, -146989073, -145645750, -144295037, -142938222,
	-141579482, -140220675, -138866091, -137517407, -136177898, -134848412,
	-133530504, -132223779, -130927987, -129641433, -128362449, -127088356,
	// 1966
	-125816717, -124544402, -123268915, -121987321, -120697707, -119398015,
	-118087401, -116765301, -115432167, -114089275, -112738221, -111381997,
	-110022780, -108664611, -107309470, -105961350, -104621292, -103292229,
	-101973810, -100667370, -99371098, -98084770, -96805351, -95531528,
	// 1967
	-94259519, -92987568, -91711771, -90430602, -89140711, -87841418,
	-86530544, -85208718, -83875380, -82532554, -81181451, -79825061,
	-78466032, -77107479, -75752735, -74404083, -73064566, -71734946,
	-70417156, -69110195, -67814577, -66527749, -65248969, -63974637,
	// 1968
	-62703242, -61430765, -60155565, -58873856, -57584541, -56284685,
	-54974349, -53651931, -52319049, -50975643, -49624847, -48267997,
	-46909093, -45550348, -44195558, -42847017, -41507302, -40178014,
	-38859920, -37553408, -36257424, -34971068, -33691882, -32417989,
	// 1969
	-31146164, -29874080, -28598437, -27317098, -26027330, -24727877,
	-23417071, -22095152, -20761777, -19418977, -18067671, -16711457,
	-15352069, -13993873, -12638718, -11290565, -9950644, -8621556,
	-7303372, -5996910, -4700895, -3414512, -2135300, -861357,
	// 1970
	410522, 1682649, 2958363, 4239725, 5529523, 6828992,
	8139714, 9461702, 10794829, 12137836, 13488728, 14845349,
	16204225, 17563004, 18917638, 20266424, 21605862, 22935526,
	24253282, 25560241, 26855849, 28142654, 29421421, 30695723,
	// 1971
	31967090, 33239546, 34514708, 35796399, 37085662, 38385465,
	39695734, 41018025, 42350858, 43694065, 45044895, 46401537,
	47760631, 49119250, 50474381, 51822888, 53162986, 54492269,
	55810689, 57117165, 58413371, 59699612, 60978917, 62252610,
	// 1972
	63524489, 64796322, 66071999, 67353072, 68642877, 69942078,
	71252926, 72574643, 73908069, 75250770, 76602122, 77958374,
	79317780, 80676161, 82031322, 83379795, 84719720, 86049179,
	87367322, 88674106, 89969985, 91256582, 92535545, 93809594,
	// 1973
	95081140, 96353315, 97628676, 98910096, 100199584, 101499179,
	102809667, 104131857, 105464820, 106808066, 108158846, 109515670,
	110874475, 112233366, 113587999, 114936836, 116276390, 117606090,
	118924056, 120231030, 121526878, 122813661, 124092645, 125366884,
	// 1974
	126638418, 127910756, 129186018, 130467528, 131756836, 133056406,
	134366710, 135688740, 137021640, 138364569, 139715503, 141071857,
	142431063, 143789404, 145144620, 146492907, 147833091, 149162293,
	150480867, 151787423, 153083864, 154370290, 155649860, 156923738,
	// 1975
	158195832, 159467756, 160743531, 162024560, 163314322, 164613372,
	165924062, 167245606, 168578804, 169921393, 171272492, 172628757,
	173987935, 175346478, 176701467, 178050189, 179389969, 180719685,
	182037702, 183344744, 184640542, 185927428, 187206353, 188480717,
	// 1976
	189752228, 191024694, 192299959, 193581587, 194870882, 196170571,
	197480787, 198802979, 200135670, 201478874, 202829482, 204186260,
	205545062, 206903917, 208258716, 209607512, 210947311, 212276911,
	213595103, 214901902, 216197937, 217484513, 218763683, 220037734,
	// 1977
	221309491, 222581693, 223857233, 225138657, 226428283, 227727773,
	229038382, 230360271, 231693399, 233036101, 234387162, 235743262,
	237102506, 238460648, 239815843, 241164043, 242504171, 243833384,
	245151866, 246458467, 247754773, 249041240, 250320666, 251594605,
	// 1978
	252866609, 254138655, 255414432, 256695669, 257985502, 259284823,
	260595569, 261917379, 263250517, 264593308, 265944185, 267300572,
	268659416, 270018012, 271372654, 272721398, 274060931, 275390712,
	276708643, 278015819, 279311629, 280598658, 281877582, 283152037,
	// 1979
	284423472, 285695974, 286971114, 288252768, 289541950, 290841689,
	292151850, 293474095, 294806803, 296150001, 297500680, 298857338,
	300216246, 301574883, 302929825, 304278375, 305618357, 306947754,
	308266176, 309572844, 310869145, 312155627, 313435051, 314708970,
	// 1980
	315980919, 317252905, 318528558, 319809692, 321099386, 322398578,
	323709283, 325030964, 326364276, 327706934, 329058237, 330414433,
	331773849, 333132133, 334487326, 335835652, 337175619, 338504934,
	339823169, 341129868, 342425914, 343712506, 344991700, 346265790,
	// 1981
	347537585, 348809787, 350085353, 351366729, 352656338, 353955803,
	355266335, 356588346, 357921322, 359264400, 360615196, 361971919,
	363330750, 364689619, 366044262, 367393117, 368732618, 370062335,
	371380195, 372687192, 373982931, 375269779, 376548696, 377823054,
	// 1982
	379094575, 380367068, 381642340, 382923998, 384213284, 385512958,
	386823169, 388145253, 389478004, 390820976, 392171755, 393528179,
	394887273, 396245720, 397600898, 398949302, 400289488, 401618754,
	402937310, 404243847, 405540224, 406826576, 408106063, 409379867,
	// 1983
	410651902, 411923793, 413199559, 414480607, 415770403, 417069494,
	418380232, 419701781, 421035024, 422377559, 423728714, 425084891,
	426444162, 427802617, 429157746, 430506418, 431846370, 433176065,
	434494237, 435801230, 437097106, 438383874, 439662796, 440936975,
	// 1984
	442208437, 443480693, 444755919, 446037369, 447326676, 448626257,
	449936539, 451258687, 452591462, 453934663, 455285325, 456642143,
	458000957, 459359903, 460714684, 462063622, 463403401, 464733185,
	466051368, 467358354, 468654349, 469941057, 471220104, 472494190,
	// 1985
	473765729, 475037877, 476313135, 477594471, 478883813, 480183256,
	481493646, 482815577, 484148582, 485491406, 486842427, 488198676,
	489557945, 490916216, 492271486, 493619771, 494960010, 496289271,
	497607895, 498914528, 500210984, 501497460, 502776995, 504050877,
	// 1986
	505322898, 506594788, 507870476, 509151464, 510441139, 511740170,
	513050774, 514372332, 515701836, 517044472, 518395457, 519751793,
	521110840, 522469457, 523824329, 525173137, 526512864, 527846318,
	529164389, 530471632, 531767543, 533054629, 534333623, 535608097,
	// 1987
	536879554, 538152000, 539427078, 540708575, 541997595, 543297097,
	544607024, 545925428, 547257908, 548600972, 549951507, 551308214,
	552667092, 554025935, 555380928, 556729762, 558069818, 559403087,
	560721554, 562028429, 563324721, 564611346, 565890716, 567164738,
	// 1988
	568436597, 569708650, 570984167, 572265310, 573554799, 574853924,
	576164354, 577482297, 578815315, 580157812, 581508902, 582865000,
	584224382, 585582675, 586938028, 588286455, 589626705, 590959747,
	592278287, 593585068, 594881358, 596167946, 597447297, 598721303,
	// 1989
	599993186, 601265249, 602540861, 603822063, 605111682, 606410931,
	607721429, 609039572, 610372471, 611715248, 613065949, 614422415,
	615781199, 617139959, 618494661, 619843600, 621183262, 622516806,
	623834869, 625142137, 626438038, 627725100, 629004079, 630278542,
	// 1990
	631550015, 632822510, 634097652, 635379251, 636668365, 637967963,
	639277980, 640596395, 641928926, 643271840, 644622374, 645978761,
	647337622, 648696081, 650051119, 651399633, 652739829, 654072910,
	655391607, 656698417, 657994988, 659281594, 660561227, 661835195,
	// 1991
	663107263, 664379199, 665654878, 666935870, 668225505, 669524486,
	670835053, 672152876, 673485987, 674828387, 676179467, 677535492,
	678894745, 680253035, 681608203, 682956739, 684296809, 685630056,
	686948441, 688255486, 689551650, 690838524, 692117743, 693392003,
	// 1992
	694663699, 695935941, 697211289, 698492604, 699781922, 701081280,
	702391508, 703713415, 705046125, 706389134, 707739745, 709096453,
	710455218, 711814134, 713168847, 714517813, 715857508, 717187377,
	718505501, 719812642, 721108638, 722395570, 723674672, 724949015,
	// 1993
	726220615, 727492993, 728768256, 730049737, 731338984, 732638470,
	733948662, 735270570, 736603333, 737946135, 739296945, 740653214,
	742012352, 743370679, 744725905, 746074243, 747414490, 748743768,
	750062421, 751369044, 752665550, 753952025, 755231643, 756505561,
	// 1994
	757777698, 759049654, 760325464, 761606506, 762896267, 764195286,
	765505912, 766827362, 768160444, 769502903, 770853883, 772210043,
	773569149, 774927648, 776282646, 777631407, 778971286, 780301131,
	781619322, 782926537, 784222508, 785509529, 786788544, 788062937,
	// 1995
	789334421, 790606803, 791881947, 793163418, 794452538, 795752039,
	797062060, 798384061, 799716574, 801059623, 802410118, 803766833,
	805125631, 806484550, 807839475, 809188460, 810528486, 811858354,
	813176810, 814483870, 815780117, 817066864, 818346118, 819620191,
	// 1996
	820891874, 822163941, 823439268, 824720441, 826009777, 827308984,
	828619320, 829940994, 831273964, 832616592, 833967653, 835323832,
	836683209, 838041534, 839396942, 840745382, 842085757, 843415220,
	844733939, 846040744, 847337217, 848623792, 849903266, 851177179,
	// 1997
	852449093, 853720977, 854996545, 856277518, 857567078, 858866112,
	860176609, 861498204, 862831202, 864173909, 865524787, 866881230,
	868240197, 869598958, 870953810, 872302781, 873642556, 874972572,
	876290730, 877598104, 878894094, 880181269, 881460309, 882734840,
	// 1998
	884006305, 885278775, 886553818, 887835299, 889124238, 890423678,
	891733501, 893055408, 894387792, 895730725, 897081201, 898437749,
	899796616, 901155309, 902510372, 903859116, 905199333, 906529011,
	907847724, 909154695, 910451281, 911738030, 913017671, 914291765,
	// 1999
	915563807, 916835819, 918111401, 919392385, 920681837, 921980723,
	923291052, 924612336, 925945237, 927287521, 928638521, 929994518,
	931353871, 932712221, 934067621, 935416237, 936756570, 938086264,
	939404879, 940711916, 942008258, 943295078, 944574437, 945848619,
	// 2000
	947120434, 948392577, 949668019, 950949197, 952238558, 953537718,
	954847923, 956169579, 957502221, 958844977, 960195527, 961552075,
	962910849, 964269777, 965624596, 966973731, 968313570, 969643675,
	970961912, 972269268, 973565305, 974852385, 976131449, 977405876,
	// 2001
	978677387, 979949810, 981224964, 982506472, 983795590, 985095082,
	986405102, 987726991, 989059530, 990402294, 991752857, 993109104,
	994468041, 995826411, 997181575, 998530060, 999870400, 1001199894,
	1002518726, 1003825557, 1005122231, 1006408843, 1007688548, 1008962504,
	// 2002
	1010234627, 1011506538, 1012782262, 1014063214, 1015352866, 1016651778,
	1017962303, 1019283630, 1020616634, 1021958938, 1023309875, 1024665853,
	1026024960, 1027383281, 1028738345, 1030087003, 1031427045, 1032756905,
	1034075342, 1035382652, 1036678889, 1037966003, 1039245233, 1040519645,
	// 2003
	1041791247, 1043063540, 1044338705, 1045619997, 1046909076, 1048208366,
	1049518329, 1050840142, 1052172602, 1053515513, 1054865948, 1056222592,
	1057581304, 1058940215, 1060295029, 1061644063, 1062983992, 1064313989,
	1065632415, 1066939690, 1068235976, 1069522986, 1070802295, 1072076614,
	// 2004
	1073348302, 1074620534, 1075895768, 1077176997, 1078466136, 1079765317,
	1081075398, 1082397026, 1083729752, 1085072360, 1086423233, 1087779421,
	1089138687, 1090497024, 1091852393, 1093200814, 1094541194, 1095870611,
	1097189382, 1098496155, 1099792740, 1101079327, 1102358962, 1103632921,
	// 2005
	1104905005, 1106176920, 1107452610, 1108733545, 1110023140, 1111322036,
	1112632488, 1113953866, 1115286802, 1116629276, 1117980143, 1119336399,
	1120695425, 1122054071, 1123409030, 1124757951, 1126097821, 1127427809,
	1128746014, 1130053356, 1131349361, 1132636512, 1133915575, 1135190108,
	// 2006
	1136461627, 1137734122, 1139009236, 1140290735, 1141579719, 1142879137,
	1144188930, 1145510766, 1146843037, 1148185889, 1149536212, 1150892741,
	1152251475, 1153610246, 1154965231, 1156314136, 1157654321, 1158984181,
	1160302861, 1161609966, 1162906467, 1164193281, 1165472784, 1166746902,
	// 2007
	1168018785, 1169290824, 1170566265, 1171847307, 1173136649, 1174435615,
	1175745850, 1177067195, 1178399997, 1179742288, 1181093195, 1182449155,
	1183808473, 1185166782, 1186522247, 1187870850, 1189211343, 1190541053,
	1191859873, 1193166911, 1194463427, 1195750180, 1197029629, 1198303655,
	// 2008
	1199575478, 1200847403, 1202122817, 1203403768, 1204693124, 1205992097,
	1207302352, 1208623873, 1209956613, 1211299264, 1212649916, 1214006374,
	1215365226, 1216724106, 1218078989, 1219428154, 1220768068, 1222098291,
	1223416622, 1224724146, 1226020263, 1227307491, 1228586568, 1229861054,
	// 2009
	1231132475, 1232404847, 1233679815, 1234961197, 1236250085, 1237549451,
	1238859260, 1240181098, 1241513486, 1242856307, 1244206779, 1245563163,
	1246922039, 1248280570, 1249635697, 1250984341, 1252324683, 1253654339,
	1254973224, 1256280227, 1257576991, 1258863769, 1260143548, 1261417621,
	// 2010
	1262689739, 1263961672, 1265237279, 1266518142, 1267807583, 1269106332,
	1270416628, 1271737786, 1273070636, 1274412828, 1275763757, 1277119697,
	1278478935, 1279837261, 1281192533, 1282541200, 1283881462, 1285211326,
	1286529974, 1287837287, 1289133730, 1290420850, 1291700279, 1292974684,
	// 2011
	1294246455, 1295518688, 1296793950, 1298075091, 1299364167, 1300663210,
	1301973085, 1303294611, 1304626960, 1305969636, 1307320008, 1308676560,
	1310035293, 1311394284, 1312749181, 1314098413, 1315438426, 1316768651,
	1318087121, 1319394597, 1320690879, 1321978055, 1323257328, 1324531791,
	// 2012
	1325803427, 1327075783, 1328350939, 1329632253, 1330921261, 1332220465,
	1333530337, 1334851930, 1336184390, 1337526944, 1338877567, 1340233742,
	1341592860, 1342951271, 1344306653, 1345655230, 1346995759, 1348325356,
	1349644321, 1350951232, 1352247979, 1353534630, 1354814361, 1356088323,
	// 2013
	1357360447, 1358632333, 1359908035, 1361188926, 1362478520, 1363777345,
	1365087778, 1366409029, 1367741924, 1369084204, 1370435035, 1371791073,
	1373150110, 1374508591, 1375863649, 1377212526, 1378552597, 1379882668,
	1381201129, 1382508607, 1383804850, 1385092104, 1386371328, 1387645875,
	// 2014
	1388917463, 1390189882, 1391465000, 1392746372, 1394035337, 1395334627,
	1396644398, 1397966130, 1399298361, 1400641137, 1401991375, 1403347866,
	1404706478, 1406065272, 1407420138, 1408769144, 1410109268, 1411439325,
	1412758030, 1414065402, 1415361980, 1416649069, 1417928624, 1419202958,
	// 2015
	1420474809, 1421746969, 1423022281, 1424303360, 1425592510, 1426891479,
	1428201519, 1429522884, 1430855529, 1432197860, 1433548662, 1434904646,
	1436263907, 1437622202, 1438977659, 1440326211, 1441666750, 1442996411,
	1444315350, 1445622386, 1446919099, 1448205900, 1449485585, 1450759666,
	// 2016
	1452031695, 1453303623, 1454579162, 1455860024, 1457149415, 1458448215,
	1459758457, 1461079775, 1462412524, 1463755003, 1465105726, 1466462067,
	1467821020, 1469179831, 1470534801, 1471883929, 1473223886, 1474554090,
	1475872425, 1477179957, 1478476085, 1479763367, 1481042491, 1482317074,
	// 2017
	1483588569, 1484861043, 1486136072, 1487417511, 1488706395, 1490005749,
	1491315467, 1492637248, 1493969489, 1495312282, 1496662621, 1498019072,
	1499377866, 1500736547, 1502091627, 1503440439, 1504780742, 1506110530,
	1507429347, 1508736418, 1510033084, 1511319893, 1512599570, 1513873688,
	// 2018
	1515145733, 1516417748, 1517693315, 1518974283, 1520263692, 1521562526,
	1522872764, 1524193946, 1525526710, 1526868863, 1528219730, 1529575617,
	1530934893, 1532293202, 1533648622, 1534997297, 1536337763, 1537667627,
	1538986461, 1540293717, 1541590273, 1542877254, 1544156720, 1545430931,
	// 2019
	1546702706, 1547974741, 1549250026, 1550531004, 1551820151, 1553119074,
	1554429054, 1555750484, 1557082934, 1558425514, 1559775953, 1561132425,
	1562491203, 1563850196, 1565205156, 1566554494, 1567894586, 1569224987,
	1570543518, 1571851166, 1573147445, 1574434718, 1575713890, 1576988350,
	// 2020
	1578259793, 1579532071, 1580806995, 1582088218, 1583377014, 1584676179,
	1585985895, 1587307535, 1588639893, 1589982568, 1591333116, 1592689431,
	1594048481, 1595407028, 1596762387, 1598111110, 1599451694, 1600781452,
	1602100532, 1603407593, 1604704462, 1605991212, 1607270999, 1608544967,
	// 2021
	1609817033, 1611088818, 1612364355, 1613645066, 1614934449, 1616233077,
	1617543336, 1618864433, 1620197261, 1621539456, 1622890355, 1624246357,
	1625605552, 1626964010, 1628319260, 1629668121, 1631008398, 1632338486,
	1633657161, 1634964686, 1636261141, 1637548438, 1638827839, 1640102373,
	// 2022
	1641374054, 1642646352, 1643921449, 1645202581, 1646491424, 1647790401,
	1649100008, 1650421447, 1651753547, 1653096143, 1654446336, 1655802819,
	1657161465, 1658520402, 1659875330, 1661224547, 1662564716, 1663895000,
	1665213726, 1666521320, 1667817905, 1669105202, 1670384746, 1671659261,
	// 2023
	1672931060, 1674203341, 1675478521, 1676759624, 1678048539, 1679347431,
	1680657150, 1681978385, 1683310694, 1684652919, 1686003468, 1687359437,
	1688718611, 1690077002, 1691432549, 1692781253, 1694121978, 1695451777,
	1696770915, 1698078034, 1699374918, 1700661746, 1701941561, 1703215629,
	// 2024
	1704487752, 1705759633, 1707035222, 1708315985, 1709605364, 1710903984,
	1712214141, 1713535192, 1714867812, 1716209979, 1717560599, 1718916666,
	1720275609, 1721634273, 1722989367, 1724338516, 1725678698, 1727009038,
	1728327615, 1729635303, 1730931621, 1732219007, 1733498239, 1734772851,
}

var solartermAlias = [...]string{
	"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
	"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
	"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
	"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

// NewSolarterm 创建节气对象
func NewSolarterm(index int64) *Solarterm {
	if !isSupported(index) {
		return nil
	}
	return &Solarterm{index}
}

// CalcSolarterm 计算节气区间
func CalcSolarterm(t *time.Time) (p, n *Solarterm) {
	var prev, next int64
	prev = 0
	next = lenTimestamp() - 1
	ts := t.Unix()
	for dowhile := true; dowhile || next-prev > 1; {
		dowhile = false
		mid := prev + int64((next-prev)/2)
		if getTimestamp(prev) <= ts && ts < getTimestamp(mid) {
			next = mid
		} else {
			prev = mid
		}
	}
	if ts == getTimestamp(prev) && prev-1 >= 0 {
		prev--
	}
	p = NewSolarterm(prev)
	n = NewSolarterm(next)
	return
}

// SpringTimestamp 获取指定年份立春的时间
func SpringTimestamp(year int64) (time int64) {
	if year < SolartermFromYear || year > SolartermToYear {
		time = 0
	} else {
		time = getTimestamp(24*(year-SolartermFromYear) + 2)
	}
	return
}

// Equals 返回两个对象是否相同
func (solarterm *Solarterm) Equals(b *Solarterm) bool {
	return solarterm.index == b.index
}

// Alias 返回节气名称(立春雨水...)
func (solarterm *Solarterm) Alias() string {
	return solartermAlias[solarterm.index%24]
}

// Index 返回节气在索引表的索引
func (solarterm *Solarterm) Index() int64 {
	return solarterm.index
}

// Order 返回节气序数(12...)
func (solarterm *Solarterm) Order() int64 {
	return utils.OrderMod(solarterm.index+24-1, 24)
}

// Timestamp 返回当前节气的时间戳
func (solarterm *Solarterm) Timestamp() int64 {
	return getTimestamp(solarterm.index)
}

// Time 根据节气时间戳获取time.Time对象
func (solarterm *Solarterm) Time() time.Time {
	return time.Unix(solarterm.Timestamp(), 0)
}

// Prev 上一个节气
func (solarterm *Solarterm) Prev() *Solarterm {
	return NewSolarterm(solarterm.index - 1)
}

// Next 下一个节气
func (solarterm *Solarterm) Next() *Solarterm {
	return NewSolarterm(solarterm.index + 1)
}

// IsToday 该节气是否为今天
func (solarterm *Solarterm) IsToday() bool {
	now := time.Now()
	return solarterm.IsInDay(&now)
}

// IsInDay 查询是否为某一天为当前的节气
func (solarterm *Solarterm) IsInDay(t *time.Time) bool {
	s := solarterm.Time()

	t1 := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.Local)
	t2 := t1.Add(24 * time.Hour)

	return t1.Unix() <= t.Unix() && t.Unix() <= t2.Unix()
}

func getTimestamp(i int64) int64 {
	return solartermTimestamp[i]
}

func lenTimestamp() int64 {
	return int64(len(solartermTimestamp))
}

func isSupported(index int64) bool {
	return 0 <= index && index < lenTimestamp()
}
