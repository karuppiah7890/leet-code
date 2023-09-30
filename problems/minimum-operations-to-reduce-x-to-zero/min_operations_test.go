package min_operations_test

import (
	"testing"

	min_operations "github.com/karuppiah7890/leet-code/problems/minimum-operations-to-reduce-x-to-zero"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums           []int
	x              int
	expectedOutput int
}

type TestCases []TestCase

func TestMinOperations(t *testing.T) {
	testCases := TestCases{
		{nums: []int{}, x: 4, expectedOutput: -1},
		{nums: []int{4}, x: 4, expectedOutput: 1},
		{nums: []int{5}, x: 4, expectedOutput: -1},
		{nums: []int{3}, x: 4, expectedOutput: -1},
		{nums: []int{5, 5}, x: 4, expectedOutput: -1},
		{nums: []int{4, 4}, x: 4, expectedOutput: 1},
		{nums: []int{5, 4}, x: 4, expectedOutput: 1},
		{nums: []int{4, 5}, x: 4, expectedOutput: 1},
		{nums: []int{4, 3, 5}, x: 4, expectedOutput: 1},
		{nums: []int{5, 3, 4}, x: 4, expectedOutput: 1},
		{nums: []int{5, 6, 7, 8, 9}, x: 4, expectedOutput: -1},
		{nums: []int{2, 2}, x: 4, expectedOutput: 2},
		{nums: []int{2, 2, 2}, x: 6, expectedOutput: 3},
		{nums: []int{10, 2, 2, 2}, x: 6, expectedOutput: 3},
		{nums: []int{2, 2, 2, 10}, x: 6, expectedOutput: 3},
		// Consider a case where there are multiple ways to make x as 0, then
		// choose the minimum number of operations to make x as 0
		{nums: []int{1, 1, 4, 2, 3}, x: 5, expectedOutput: 2}, // 1,1,3 or 3,2
		{nums: []int{3, 2, 4, 1, 1}, x: 5, expectedOutput: 2}, // 1,1,3 or 3,2
		{nums: []int{3, 2, 20, 1, 1, 3}, x: 10, expectedOutput: 5},
		{nums: []int{5, 2, 5}, x: 6, expectedOutput: -1},
		{nums: []int{1241, 8769, 9151, 3211, 2314,
			8007, 3713, 5835, 2176, 8227, 5251,
			9229, 904, 1899, 5513, 7878, 8663,
			3804, 2685, 3501, 1204, 9742, 2578,
			8849, 1120, 4687, 5902, 9929, 6769,
			8171, 5150, 1343, 9619, 3973, 3273,
			6427, 47, 8701, 2741, 7402, 1412,
			2223, 8152, 805, 6726, 9128, 2794,
			7137, 6725, 4279, 7200, 5582, 9583,
			7443, 6573, 7221, 1423, 4859, 2608,
			3772, 7437, 2581, 975, 3893, 9172,
			3, 3113, 2978, 9300, 6029, 4958, 229,
			4630, 653, 1421, 5512, 5392, 7287, 8643,
			4495, 2640, 8047, 7268, 3878, 6010, 8070,
			7560, 8931, 76, 6502, 5952, 4871, 5986,
			4935, 3015, 8263, 7497, 8153, 384, 1136}, x: 894887480, expectedOutput: -1},

		{nums: []int{1072, 2718, 6852, 2116, 439, 442, 8829, 6511, 1461, 6872, 3721, 5661, 8151, 1730, 9212, 9920, 3817, 8888, 6674, 747, 2928, 8334, 6371, 4397, 5676, 500, 4510, 2082, 1472, 5599, 8219, 1359, 4749, 2192, 87, 2814, 3819, 731, 5180, 6840, 4825, 2330, 3287, 9906, 8437, 8260, 9953, 3067, 9604, 8415, 6475, 4867, 8768, 536, 3186, 5412, 4410, 5161, 6458, 5478, 929, 7907, 1679, 3216, 879, 8154, 926, 3663, 2052, 8392, 1506, 6230, 8095, 9647, 5045, 5837, 4040, 4940, 2317, 6932, 9500, 5837, 2568, 5344, 4460, 2602, 1040, 6750, 7334, 6627, 2572, 1908, 4962, 8496, 4087, 1611, 1154, 6854, 7428, 4558, 614, 3584, 1318, 9310, 4181, 145, 1403, 3507, 4920, 4926, 7094, 2037, 4545, 9078, 9813, 944, 3531, 2685, 5414, 3217, 8442, 8007, 5221, 3838, 8850, 6115, 239, 6489, 7842, 9418, 3413, 1114, 6620, 3860, 860, 5191, 4945, 6594, 5046, 3897, 8833, 3191, 3698, 4434, 2875, 3660, 9451, 3343, 6588, 4510, 9912, 4100, 60, 6913, 1267, 6189, 8737, 4772, 2063, 2434, 9354, 9212, 4959, 8422, 6028, 3374, 7594, 3336, 5351, 9747, 3497, 1425, 3159, 7293, 1282, 4198, 5701, 5233, 9388, 480, 315, 8938, 7514, 1352, 5535, 5706, 7596, 7486, 1059, 1164, 1857, 6155, 2026, 5483, 2499, 4768, 3083, 7514, 2076, 7853, 9274, 2107, 6430, 6743, 280, 8774, 7505, 5545, 4919, 7515, 7892, 2792, 1120, 4859, 5032, 4735, 4430, 6868, 8662, 4961, 5834, 5355, 1166, 1352, 136, 2008, 1287, 4253, 8651, 7525, 6826, 2360, 9110, 5732, 4169, 9682, 7677, 8278, 6331, 8356, 5958, 8900, 2869, 847, 9130, 8136, 4138, 9544, 6055, 7514, 3162, 1945, 3957, 8876, 6743, 6909, 3895, 6258, 4003, 1269, 6714, 7060, 307, 1180, 6675, 7313, 845, 1054, 1224, 7381, 5648, 165, 1525, 5408, 4881, 8716, 9531, 7780, 788, 5962, 1569, 4152, 3122, 6225, 589, 7990, 1029, 5521, 4672, 3093, 8458, 4650, 8021, 9422, 2317, 3535, 7259, 5019, 88, 4705, 2804, 1211, 2470, 6478, 927, 4022, 9069, 8575, 5183, 538, 6884, 3293, 6270, 7506, 4673, 9849, 5509, 4062, 1865, 7698, 2966, 5558, 2649, 6955, 153, 5041, 9349, 5461, 4392, 5754, 6347, 5329, 6657, 5854, 9589, 6559, 5226, 2367, 5798, 5655, 6852, 1869, 8887, 4641, 3904, 7940, 4106, 1794, 4151, 5517, 6276, 4317, 1095, 6671, 2250, 7048, 2567, 7229, 3282, 7526, 7033, 4801, 3451, 7094, 5461, 859, 9315, 2576, 8259, 772, 3492, 3243, 8679, 2028, 3774, 9067, 3676, 8514, 4978, 7963, 7343, 8691, 5585, 369, 9406, 8646, 8954, 2960, 1510, 4084, 919, 1519, 8251, 2086, 2639, 3251, 8521, 5876, 8871, 1474, 7022, 3703, 2092, 2092, 2465, 7793, 5924, 4219, 1615, 5693, 1581, 2695, 3128, 8979, 5859, 7066, 3721, 1483, 6458, 878, 6334, 2951, 4479, 4603, 836, 7063, 8897, 5412, 341, 180, 9936, 9546, 9110, 1365, 4663, 6152, 1223, 1642, 3302, 4337, 4906, 6365, 8761, 6084, 1864, 4935, 3603, 9965, 2492, 4531, 6253, 8466, 319, 1085, 2419, 9284, 7619, 7954, 5357, 8381, 9624, 3433, 746, 4733, 1310, 6688, 9752, 4073, 2877, 2524, 6968, 7959, 5084, 1790, 47, 6790, 7447, 3739, 4870, 7222, 8208, 6715, 1900, 4439, 96, 8346, 4078, 511, 2719, 9021, 9556, 1208, 3075, 1759, 2803, 1536, 9453, 1118, 3537, 8028, 3524, 4974, 7366, 681, 5076, 7245, 4189, 3230, 2224, 1919, 1977, 3628, 9610, 2716, 8476, 8039, 8766, 7163, 1029, 2866, 5892, 102, 1461, 8308, 454, 5822, 7555, 2903, 8470, 5973, 246, 8652, 9846, 7901, 2708, 3973, 9743, 3466, 7611, 8645, 5904, 6254, 1309, 7198, 110, 5938, 4301, 9478, 1298, 7589, 8462, 2583, 8067, 3947, 8196, 8037, 4333, 9432, 9601, 4121, 1263, 3920, 4788, 7251, 8145, 2472, 6012, 3598, 6649, 9734, 5670, 6030, 3789, 73, 9455, 3971, 7434, 5948, 7593, 691, 4564, 6814, 3089, 3188, 3564, 5559, 6746, 4458, 4078, 2311, 4832, 5510, 6001, 5341, 1194, 4668, 6746, 7273, 8605, 1632, 9711, 7436, 8968, 1022, 9032, 6984, 9092, 3221, 9238, 1467, 3168, 5688, 4613, 37, 7567, 4485, 1296, 8952, 8263, 1313, 2823, 3017, 4140, 8309, 1689, 4699, 3281, 3882, 1777, 5914, 9561, 2160, 9996, 3302, 9569, 2684, 791, 2348, 5919, 6198, 4526, 3823, 120, 6719, 4657, 8076, 2051, 535, 7275, 4314, 6977, 9380, 2348, 3608, 1701, 6090, 3133, 3023, 7486, 4462, 5157, 6185, 8389, 5376, 6752, 770, 1873, 9493, 9635, 7886, 5416, 677, 2362, 5052, 3594, 8195, 64, 9757, 9515, 9701, 7987, 4681, 5397, 2608, 3031, 6801, 5637, 3806, 9843, 9665, 5443, 5313, 3772, 8335, 7437, 248, 4617, 6558, 6953, 6185, 5465, 1472, 4749, 4862, 4112, 8928, 9513, 8568, 6408, 8414, 4184, 8075, 137, 1725, 6546, 9235, 2017, 4311, 9415, 4912, 2849, 5863, 4888, 411, 910, 269, 7851, 3645, 1184, 9437, 8078, 1801, 2505, 8795, 9233, 614, 1332, 7945, 350, 8501, 4367, 706, 2776, 9794, 792, 5065, 2892, 6781, 823, 8522, 2898, 4028, 924, 1897, 1674, 8692, 6154, 9217, 7890, 9632, 7322, 5088, 1231, 8999, 2685, 5681, 9058, 5126, 699, 752, 7029, 5495, 6858, 3141, 7101, 4184, 9645, 3390, 9424, 8373, 7610, 1722, 2193, 8292, 4061, 2814, 3706, 9238, 7540, 3738, 5273, 2169, 6902, 3387, 5420, 7762, 7031, 3163, 8064, 7294, 160, 8346, 8542, 2350, 7038, 4100, 8896, 9296, 5612, 1955, 9934, 2833, 3492, 9752, 2436, 3554, 6166, 9718, 747, 1286, 8056, 3363, 4194, 6456, 2111, 734, 9929, 1005, 3218, 940, 6137, 4586, 2712, 9347, 2670, 3216, 1830, 9482, 8439, 3071, 7471, 4658, 159, 9964, 9219, 959, 7566, 7737, 1106, 1966, 1281, 2433, 487, 2491, 212, 4719, 1446, 2123, 5254, 8868, 8198, 1135, 5760, 1550, 2868, 8064, 496, 6335, 9900, 1422, 1657, 748, 5153, 295, 8789, 8706, 9716, 5380, 2716, 4628, 6654, 3240, 4033, 4851, 1675, 6660, 3107, 4960, 7307, 429, 8986, 7936, 6337, 8786, 5136, 3431, 8132, 1328, 6924, 6911, 1353, 2600, 6786, 5575, 3394, 3659, 9497, 8444, 7654, 5573, 2822, 64, 5975, 893, 7968, 7222, 1192, 3976, 6229, 7658, 71, 7911, 2165, 1470, 4924, 7512, 8327, 8823, 2582, 2271, 2140, 5991, 6342, 2546, 6175, 9732, 9273, 6647, 5853, 8451, 4805, 8274, 1125, 7909, 5517, 7300, 3136, 689, 9377, 3589, 6299, 5330, 9543, 5925, 8329, 1157, 8548, 9794, 5919, 8068, 8044, 9233, 5737, 5142, 3460, 5932, 6523, 9508, 1957, 8332, 342, 9217, 945, 9604, 8404, 1, 1436, 4046, 6616, 8369, 3723, 208, 7708, 1048, 5625, 1639, 4658, 3339, 1497, 7821, 1868, 6336, 1280, 858, 1044, 3349, 7431, 7778, 2218, 8886, 5086, 2375, 6367, 5760, 2980, 8930, 7419, 5134, 8355, 8575, 1904, 9316, 3544, 7774, 4203, 6460, 197, 4286, 5925, 4934, 676, 9488, 7870, 6183, 9245, 6625, 2128, 645, 4983, 3084, 117, 984, 4955, 1771, 8856, 9201, 4929, 3002, 6430, 3085, 1420, 5269, 6, 4508, 7679, 2251, 3932, 5693, 983, 4194, 6727, 9944, 6786, 9147, 6688, 3044, 4334, 8832, 3794, 3667, 5779, 1004, 4982, 560, 1262, 7415, 8759, 7812, 579, 917, 9967, 4392, 9850, 1359, 3604, 8676, 4700, 5733, 4145, 2152, 358, 2109, 2242, 8400, 2983, 4728, 4346, 386, 9743, 9459, 8122, 8512, 628, 9349, 1038, 2931, 8209, 470, 8854, 8387, 4909, 2947, 3286, 1760, 1037, 7367, 8512, 7175, 2163, 2699, 5617, 4068, 173, 1975, 5478, 9569, 2970, 1176, 953, 2882, 458, 9678, 7350, 915, 3677, 4148, 9686, 3400, 6532, 2013, 6401, 8290, 54, 8253, 9291, 7235, 8186, 5308, 3266, 8245, 3643, 8800, 6543, 8251, 4472, 486, 5046, 152, 3250, 7702, 701, 6725, 9608, 8491, 5145, 2403, 8885, 6539, 6713, 8438, 9200, 907, 3108, 9975, 768, 8773, 8682, 669, 5107, 4243, 5328, 733, 2618, 5503, 9459, 6720, 8058, 314, 6433, 5321, 3788, 1228, 4810, 986, 9639, 2554, 9367, 26, 5164, 7273, 2052, 2541, 721, 467, 7802, 3398, 8497, 7227, 8795, 4042, 9451, 8086, 1824, 7656, 3354, 7299, 196, 8020, 2752, 4975, 9496, 9576, 4339, 7474, 7680, 1667, 4071, 3150, 7894, 7410, 868, 7425, 6214, 2648, 6774, 5782, 9081, 8495, 5292, 5536, 6730, 5330, 5424, 8151, 2975, 9050, 6790, 4751, 4021, 8832, 6824, 4844, 7478, 2670, 1635, 9774, 3518, 6818, 2295, 3305, 9335, 432, 7464, 3533, 8936, 3823, 5585, 8197, 5689, 4432, 8443, 8021, 2277, 3641, 5921, 5191, 7366, 5862, 4736, 3864, 450, 5737, 4495, 4897, 25, 1854, 7040, 908, 8996, 4744, 4803, 3826, 734, 8344, 223, 3686, 5679, 9536, 267, 2910, 3233, 8768, 2326, 4574, 1601, 6971, 5600, 508, 8023, 8112, 3437, 7944, 1228, 3406, 4099, 1021, 9717, 3977, 7216, 3323, 214, 5425, 5471, 1012, 4060, 9596, 5299, 7996, 1112, 178, 2992, 801, 9244, 2871, 3420, 728, 5849, 3497, 1100, 2989, 9041, 2231, 6853, 3623, 6467, 7381, 3608, 1127, 8613, 1951, 6645, 7083, 9468, 9625, 3005, 5484, 4148, 6126, 1887, 7805, 915, 1362, 9149, 7527, 1333, 293, 6769, 8566, 9258, 4576, 7961, 4033, 4709, 2761, 9522, 6848, 662, 7470, 5642, 2306, 4753, 4827, 7358, 9470, 8957, 5999, 4517, 1548, 6282, 793, 2641, 6510, 7834, 5718, 7666, 2857, 2929, 5955, 8145, 8588, 8048, 2708, 6307, 821, 9315, 4914, 6015, 8624, 5875, 9992, 9145, 7405, 90, 4121, 9835, 5771, 4406, 2721, 1744, 1127, 5059, 890, 1293, 3319, 4458, 9091, 1380, 7430, 5471, 5217, 8787, 2389, 7012, 2769, 1163, 9089, 7424, 3149, 1753, 6317, 4553, 3924, 3367, 6018, 8683, 577, 9837, 3052, 804, 870, 2024, 6176, 5188, 5628, 2146, 3896, 3783, 8487, 7764, 4658, 8209, 4353, 1204, 7136, 1507, 1741, 3585, 8041, 202, 5336, 9791, 9658, 808, 857, 2265, 2187, 4701, 8443, 1320, 5562, 8255, 6870, 7267, 2976, 520, 5913, 9092, 8912, 8333, 2001, 9005, 3390, 3978, 8258, 7959, 4442, 7343, 8195, 6831, 1789, 7905, 7405, 3933, 9069, 3299, 1028, 6846, 1298, 4321, 1518, 9865, 806, 6195, 9570, 3592, 2803, 5439, 9341, 364, 5770, 3668, 3623, 6636, 4226, 604, 6332, 6702, 8654, 3729, 5103, 6908, 5993, 7624, 2415, 6245, 8508, 1015, 7780, 3868, 6405, 3258, 4993, 1404, 2386, 688, 5982, 9553, 5997, 5185, 8280, 2982, 155, 9556, 3693, 211, 10000, 7974, 8216, 8576, 37, 4150, 1280, 5560, 4277, 5805, 9740, 3254, 7775, 7974, 2311, 5112, 8448, 480, 514, 4849, 8288, 6919, 8885, 2297, 2626, 8779, 7746, 5373, 7070, 6083, 8161, 320, 700, 5885, 651, 8259, 7231, 5228, 2450, 6196, 6989, 7046, 2967, 2468, 5489, 3214, 465, 3838, 6064, 4100, 6030, 9416, 6048, 1286, 8120, 8524, 7313, 3392, 8920, 3671, 653, 8803, 6786, 2024, 4190, 5365, 8355, 1682, 3610, 2139, 4641, 5772, 3575, 3866, 7468, 9214, 4425, 4843, 1438, 6294, 1106, 7035, 5643, 2403, 6586, 3903, 7651, 1968, 151, 4207, 5847, 9762, 613, 6651, 8481, 8912, 2971, 4170, 9201, 9646, 5669, 7798, 1733, 8860, 3446, 126, 4089, 1671, 3499, 5330, 6828, 5732, 4753, 5530, 446, 5791, 2669, 2894, 6495, 9068, 1364, 2899, 94, 8585, 634, 4919, 2667, 8001, 8334, 4645, 2976, 5013, 5876, 463, 6896, 914, 5405, 5626, 3841, 8835, 5100, 9244, 907, 2614, 5965, 2695, 5263, 8446, 1262, 1248, 4262, 1706, 2511, 2302, 8568, 7679, 8917, 3856, 8929, 9366, 3885, 9032, 9654, 699, 6391, 1491, 740, 9239, 1185, 5629, 6287, 2056, 287, 8348, 3534, 6644, 6918, 5285, 714, 5609, 6986, 9304, 9020, 2867, 7001, 4794, 4706, 4232, 6790, 9242, 4740, 5339, 2899, 9572, 584, 4638, 6185, 4907, 5092, 7729, 8613, 5310, 624, 9112, 4506, 4329, 5427, 8372, 2412, 3736, 862, 7833, 9342, 1775, 8568, 6089, 6928, 206, 6083, 4158, 1939, 4986, 8806, 5452, 5212, 1634, 3942, 9004, 6675, 5185, 1459, 6796, 592, 6782, 8691, 3634, 7161, 8748, 559, 7607, 661, 1082, 2762, 8716, 9638, 6770, 5767, 3236, 8591, 8404, 1449, 8995, 6510, 3133, 6456, 4567, 4809, 7712, 131, 6707, 9003, 5730, 6567, 6693, 3642, 6477, 5833, 5796, 5699, 9603, 2938, 2192, 1687, 7398, 9145, 2423, 4412, 246, 8040, 6030, 7415, 9271, 8094, 8456, 7274, 5757, 6014, 6731, 5728, 4218, 3875, 7103, 5833},
			x:              7327176,
			expectedOutput: 1450,
		},
	}

	for _, testCase := range testCases {
		output := min_operations.MinOperations(testCase.nums, testCase.x)

		assert.Equal(t, testCase.expectedOutput, output, "Input: nums: %v, x: %v", testCase.nums, testCase.x)
	}
}

func BenchmarkMinOperations(b *testing.B) {
	nums := []int{1072, 2718, 6852, 2116, 439, 442, 8829, 6511, 1461, 6872, 3721, 5661, 8151, 1730, 9212, 9920, 3817, 8888, 6674, 747, 2928, 8334, 6371, 4397, 5676, 500, 4510, 2082, 1472, 5599, 8219, 1359, 4749, 2192, 87, 2814, 3819, 731, 5180, 6840, 4825, 2330, 3287, 9906, 8437, 8260, 9953, 3067, 9604, 8415, 6475, 4867, 8768, 536, 3186, 5412, 4410, 5161, 6458, 5478, 929, 7907, 1679, 3216, 879, 8154, 926, 3663, 2052, 8392, 1506, 6230, 8095, 9647, 5045, 5837, 4040, 4940, 2317, 6932, 9500, 5837, 2568, 5344, 4460, 2602, 1040, 6750, 7334, 6627, 2572, 1908, 4962, 8496, 4087, 1611, 1154, 6854, 7428, 4558, 614, 3584, 1318, 9310, 4181, 145, 1403, 3507, 4920, 4926, 7094, 2037, 4545, 9078, 9813, 944, 3531, 2685, 5414, 3217, 8442, 8007, 5221, 3838, 8850, 6115, 239, 6489, 7842, 9418, 3413, 1114, 6620, 3860, 860, 5191, 4945, 6594, 5046, 3897, 8833, 3191, 3698, 4434, 2875, 3660, 9451, 3343, 6588, 4510, 9912, 4100, 60, 6913, 1267, 6189, 8737, 4772, 2063, 2434, 9354, 9212, 4959, 8422, 6028, 3374, 7594, 3336, 5351, 9747, 3497, 1425, 3159, 7293, 1282, 4198, 5701, 5233, 9388, 480, 315, 8938, 7514, 1352, 5535, 5706, 7596, 7486, 1059, 1164, 1857, 6155, 2026, 5483, 2499, 4768, 3083, 7514, 2076, 7853, 9274, 2107, 6430, 6743, 280, 8774, 7505, 5545, 4919, 7515, 7892, 2792, 1120, 4859, 5032, 4735, 4430, 6868, 8662, 4961, 5834, 5355, 1166, 1352, 136, 2008, 1287, 4253, 8651, 7525, 6826, 2360, 9110, 5732, 4169, 9682, 7677, 8278, 6331, 8356, 5958, 8900, 2869, 847, 9130, 8136, 4138, 9544, 6055, 7514, 3162, 1945, 3957, 8876, 6743, 6909, 3895, 6258, 4003, 1269, 6714, 7060, 307, 1180, 6675, 7313, 845, 1054, 1224, 7381, 5648, 165, 1525, 5408, 4881, 8716, 9531, 7780, 788, 5962, 1569, 4152, 3122, 6225, 589, 7990, 1029, 5521, 4672, 3093, 8458, 4650, 8021, 9422, 2317, 3535, 7259, 5019, 88, 4705, 2804, 1211, 2470, 6478, 927, 4022, 9069, 8575, 5183, 538, 6884, 3293, 6270, 7506, 4673, 9849, 5509, 4062, 1865, 7698, 2966, 5558, 2649, 6955, 153, 5041, 9349, 5461, 4392, 5754, 6347, 5329, 6657, 5854, 9589, 6559, 5226, 2367, 5798, 5655, 6852, 1869, 8887, 4641, 3904, 7940, 4106, 1794, 4151, 5517, 6276, 4317, 1095, 6671, 2250, 7048, 2567, 7229, 3282, 7526, 7033, 4801, 3451, 7094, 5461, 859, 9315, 2576, 8259, 772, 3492, 3243, 8679, 2028, 3774, 9067, 3676, 8514, 4978, 7963, 7343, 8691, 5585, 369, 9406, 8646, 8954, 2960, 1510, 4084, 919, 1519, 8251, 2086, 2639, 3251, 8521, 5876, 8871, 1474, 7022, 3703, 2092, 2092, 2465, 7793, 5924, 4219, 1615, 5693, 1581, 2695, 3128, 8979, 5859, 7066, 3721, 1483, 6458, 878, 6334, 2951, 4479, 4603, 836, 7063, 8897, 5412, 341, 180, 9936, 9546, 9110, 1365, 4663, 6152, 1223, 1642, 3302, 4337, 4906, 6365, 8761, 6084, 1864, 4935, 3603, 9965, 2492, 4531, 6253, 8466, 319, 1085, 2419, 9284, 7619, 7954, 5357, 8381, 9624, 3433, 746, 4733, 1310, 6688, 9752, 4073, 2877, 2524, 6968, 7959, 5084, 1790, 47, 6790, 7447, 3739, 4870, 7222, 8208, 6715, 1900, 4439, 96, 8346, 4078, 511, 2719, 9021, 9556, 1208, 3075, 1759, 2803, 1536, 9453, 1118, 3537, 8028, 3524, 4974, 7366, 681, 5076, 7245, 4189, 3230, 2224, 1919, 1977, 3628, 9610, 2716, 8476, 8039, 8766, 7163, 1029, 2866, 5892, 102, 1461, 8308, 454, 5822, 7555, 2903, 8470, 5973, 246, 8652, 9846, 7901, 2708, 3973, 9743, 3466, 7611, 8645, 5904, 6254, 1309, 7198, 110, 5938, 4301, 9478, 1298, 7589, 8462, 2583, 8067, 3947, 8196, 8037, 4333, 9432, 9601, 4121, 1263, 3920, 4788, 7251, 8145, 2472, 6012, 3598, 6649, 9734, 5670, 6030, 3789, 73, 9455, 3971, 7434, 5948, 7593, 691, 4564, 6814, 3089, 3188, 3564, 5559, 6746, 4458, 4078, 2311, 4832, 5510, 6001, 5341, 1194, 4668, 6746, 7273, 8605, 1632, 9711, 7436, 8968, 1022, 9032, 6984, 9092, 3221, 9238, 1467, 3168, 5688, 4613, 37, 7567, 4485, 1296, 8952, 8263, 1313, 2823, 3017, 4140, 8309, 1689, 4699, 3281, 3882, 1777, 5914, 9561, 2160, 9996, 3302, 9569, 2684, 791, 2348, 5919, 6198, 4526, 3823, 120, 6719, 4657, 8076, 2051, 535, 7275, 4314, 6977, 9380, 2348, 3608, 1701, 6090, 3133, 3023, 7486, 4462, 5157, 6185, 8389, 5376, 6752, 770, 1873, 9493, 9635, 7886, 5416, 677, 2362, 5052, 3594, 8195, 64, 9757, 9515, 9701, 7987, 4681, 5397, 2608, 3031, 6801, 5637, 3806, 9843, 9665, 5443, 5313, 3772, 8335, 7437, 248, 4617, 6558, 6953, 6185, 5465, 1472, 4749, 4862, 4112, 8928, 9513, 8568, 6408, 8414, 4184, 8075, 137, 1725, 6546, 9235, 2017, 4311, 9415, 4912, 2849, 5863, 4888, 411, 910, 269, 7851, 3645, 1184, 9437, 8078, 1801, 2505, 8795, 9233, 614, 1332, 7945, 350, 8501, 4367, 706, 2776, 9794, 792, 5065, 2892, 6781, 823, 8522, 2898, 4028, 924, 1897, 1674, 8692, 6154, 9217, 7890, 9632, 7322, 5088, 1231, 8999, 2685, 5681, 9058, 5126, 699, 752, 7029, 5495, 6858, 3141, 7101, 4184, 9645, 3390, 9424, 8373, 7610, 1722, 2193, 8292, 4061, 2814, 3706, 9238, 7540, 3738, 5273, 2169, 6902, 3387, 5420, 7762, 7031, 3163, 8064, 7294, 160, 8346, 8542, 2350, 7038, 4100, 8896, 9296, 5612, 1955, 9934, 2833, 3492, 9752, 2436, 3554, 6166, 9718, 747, 1286, 8056, 3363, 4194, 6456, 2111, 734, 9929, 1005, 3218, 940, 6137, 4586, 2712, 9347, 2670, 3216, 1830, 9482, 8439, 3071, 7471, 4658, 159, 9964, 9219, 959, 7566, 7737, 1106, 1966, 1281, 2433, 487, 2491, 212, 4719, 1446, 2123, 5254, 8868, 8198, 1135, 5760, 1550, 2868, 8064, 496, 6335, 9900, 1422, 1657, 748, 5153, 295, 8789, 8706, 9716, 5380, 2716, 4628, 6654, 3240, 4033, 4851, 1675, 6660, 3107, 4960, 7307, 429, 8986, 7936, 6337, 8786, 5136, 3431, 8132, 1328, 6924, 6911, 1353, 2600, 6786, 5575, 3394, 3659, 9497, 8444, 7654, 5573, 2822, 64, 5975, 893, 7968, 7222, 1192, 3976, 6229, 7658, 71, 7911, 2165, 1470, 4924, 7512, 8327, 8823, 2582, 2271, 2140, 5991, 6342, 2546, 6175, 9732, 9273, 6647, 5853, 8451, 4805, 8274, 1125, 7909, 5517, 7300, 3136, 689, 9377, 3589, 6299, 5330, 9543, 5925, 8329, 1157, 8548, 9794, 5919, 8068, 8044, 9233, 5737, 5142, 3460, 5932, 6523, 9508, 1957, 8332, 342, 9217, 945, 9604, 8404, 1, 1436, 4046, 6616, 8369, 3723, 208, 7708, 1048, 5625, 1639, 4658, 3339, 1497, 7821, 1868, 6336, 1280, 858, 1044, 3349, 7431, 7778, 2218, 8886, 5086, 2375, 6367, 5760, 2980, 8930, 7419, 5134, 8355, 8575, 1904, 9316, 3544, 7774, 4203, 6460, 197, 4286, 5925, 4934, 676, 9488, 7870, 6183, 9245, 6625, 2128, 645, 4983, 3084, 117, 984, 4955, 1771, 8856, 9201, 4929, 3002, 6430, 3085, 1420, 5269, 6, 4508, 7679, 2251, 3932, 5693, 983, 4194, 6727, 9944, 6786, 9147, 6688, 3044, 4334, 8832, 3794, 3667, 5779, 1004, 4982, 560, 1262, 7415, 8759, 7812, 579, 917, 9967, 4392, 9850, 1359, 3604, 8676, 4700, 5733, 4145, 2152, 358, 2109, 2242, 8400, 2983, 4728, 4346, 386, 9743, 9459, 8122, 8512, 628, 9349, 1038, 2931, 8209, 470, 8854, 8387, 4909, 2947, 3286, 1760, 1037, 7367, 8512, 7175, 2163, 2699, 5617, 4068, 173, 1975, 5478, 9569, 2970, 1176, 953, 2882, 458, 9678, 7350, 915, 3677, 4148, 9686, 3400, 6532, 2013, 6401, 8290, 54, 8253, 9291, 7235, 8186, 5308, 3266, 8245, 3643, 8800, 6543, 8251, 4472, 486, 5046, 152, 3250, 7702, 701, 6725, 9608, 8491, 5145, 2403, 8885, 6539, 6713, 8438, 9200, 907, 3108, 9975, 768, 8773, 8682, 669, 5107, 4243, 5328, 733, 2618, 5503, 9459, 6720, 8058, 314, 6433, 5321, 3788, 1228, 4810, 986, 9639, 2554, 9367, 26, 5164, 7273, 2052, 2541, 721, 467, 7802, 3398, 8497, 7227, 8795, 4042, 9451, 8086, 1824, 7656, 3354, 7299, 196, 8020, 2752, 4975, 9496, 9576, 4339, 7474, 7680, 1667, 4071, 3150, 7894, 7410, 868, 7425, 6214, 2648, 6774, 5782, 9081, 8495, 5292, 5536, 6730, 5330, 5424, 8151, 2975, 9050, 6790, 4751, 4021, 8832, 6824, 4844, 7478, 2670, 1635, 9774, 3518, 6818, 2295, 3305, 9335, 432, 7464, 3533, 8936, 3823, 5585, 8197, 5689, 4432, 8443, 8021, 2277, 3641, 5921, 5191, 7366, 5862, 4736, 3864, 450, 5737, 4495, 4897, 25, 1854, 7040, 908, 8996, 4744, 4803, 3826, 734, 8344, 223, 3686, 5679, 9536, 267, 2910, 3233, 8768, 2326, 4574, 1601, 6971, 5600, 508, 8023, 8112, 3437, 7944, 1228, 3406, 4099, 1021, 9717, 3977, 7216, 3323, 214, 5425, 5471, 1012, 4060, 9596, 5299, 7996, 1112, 178, 2992, 801, 9244, 2871, 3420, 728, 5849, 3497, 1100, 2989, 9041, 2231, 6853, 3623, 6467, 7381, 3608, 1127, 8613, 1951, 6645, 7083, 9468, 9625, 3005, 5484, 4148, 6126, 1887, 7805, 915, 1362, 9149, 7527, 1333, 293, 6769, 8566, 9258, 4576, 7961, 4033, 4709, 2761, 9522, 6848, 662, 7470, 5642, 2306, 4753, 4827, 7358, 9470, 8957, 5999, 4517, 1548, 6282, 793, 2641, 6510, 7834, 5718, 7666, 2857, 2929, 5955, 8145, 8588, 8048, 2708, 6307, 821, 9315, 4914, 6015, 8624, 5875, 9992, 9145, 7405, 90, 4121, 9835, 5771, 4406, 2721, 1744, 1127, 5059, 890, 1293, 3319, 4458, 9091, 1380, 7430, 5471, 5217, 8787, 2389, 7012, 2769, 1163, 9089, 7424, 3149, 1753, 6317, 4553, 3924, 3367, 6018, 8683, 577, 9837, 3052, 804, 870, 2024, 6176, 5188, 5628, 2146, 3896, 3783, 8487, 7764, 4658, 8209, 4353, 1204, 7136, 1507, 1741, 3585, 8041, 202, 5336, 9791, 9658, 808, 857, 2265, 2187, 4701, 8443, 1320, 5562, 8255, 6870, 7267, 2976, 520, 5913, 9092, 8912, 8333, 2001, 9005, 3390, 3978, 8258, 7959, 4442, 7343, 8195, 6831, 1789, 7905, 7405, 3933, 9069, 3299, 1028, 6846, 1298, 4321, 1518, 9865, 806, 6195, 9570, 3592, 2803, 5439, 9341, 364, 5770, 3668, 3623, 6636, 4226, 604, 6332, 6702, 8654, 3729, 5103, 6908, 5993, 7624, 2415, 6245, 8508, 1015, 7780, 3868, 6405, 3258, 4993, 1404, 2386, 688, 5982, 9553, 5997, 5185, 8280, 2982, 155, 9556, 3693, 211, 10000, 7974, 8216, 8576, 37, 4150, 1280, 5560, 4277, 5805, 9740, 3254, 7775, 7974, 2311, 5112, 8448, 480, 514, 4849, 8288, 6919, 8885, 2297, 2626, 8779, 7746, 5373, 7070, 6083, 8161, 320, 700, 5885, 651, 8259, 7231, 5228, 2450, 6196, 6989, 7046, 2967, 2468, 5489, 3214, 465, 3838, 6064, 4100, 6030, 9416, 6048, 1286, 8120, 8524, 7313, 3392, 8920, 3671, 653, 8803, 6786, 2024, 4190, 5365, 8355, 1682, 3610, 2139, 4641, 5772, 3575, 3866, 7468, 9214, 4425, 4843, 1438, 6294, 1106, 7035, 5643, 2403, 6586, 3903, 7651, 1968, 151, 4207, 5847, 9762, 613, 6651, 8481, 8912, 2971, 4170, 9201, 9646, 5669, 7798, 1733, 8860, 3446, 126, 4089, 1671, 3499, 5330, 6828, 5732, 4753, 5530, 446, 5791, 2669, 2894, 6495, 9068, 1364, 2899, 94, 8585, 634, 4919, 2667, 8001, 8334, 4645, 2976, 5013, 5876, 463, 6896, 914, 5405, 5626, 3841, 8835, 5100, 9244, 907, 2614, 5965, 2695, 5263, 8446, 1262, 1248, 4262, 1706, 2511, 2302, 8568, 7679, 8917, 3856, 8929, 9366, 3885, 9032, 9654, 699, 6391, 1491, 740, 9239, 1185, 5629, 6287, 2056, 287, 8348, 3534, 6644, 6918, 5285, 714, 5609, 6986, 9304, 9020, 2867, 7001, 4794, 4706, 4232, 6790, 9242, 4740, 5339, 2899, 9572, 584, 4638, 6185, 4907, 5092, 7729, 8613, 5310, 624, 9112, 4506, 4329, 5427, 8372, 2412, 3736, 862, 7833, 9342, 1775, 8568, 6089, 6928, 206, 6083, 4158, 1939, 4986, 8806, 5452, 5212, 1634, 3942, 9004, 6675, 5185, 1459, 6796, 592, 6782, 8691, 3634, 7161, 8748, 559, 7607, 661, 1082, 2762, 8716, 9638, 6770, 5767, 3236, 8591, 8404, 1449, 8995, 6510, 3133, 6456, 4567, 4809, 7712, 131, 6707, 9003, 5730, 6567, 6693, 3642, 6477, 5833, 5796, 5699, 9603, 2938, 2192, 1687, 7398, 9145, 2423, 4412, 246, 8040, 6030, 7415, 9271, 8094, 8456, 7274, 5757, 6014, 6731, 5728, 4218, 3875, 7103, 5833}
	x := 7327176

	for i := 0; i < b.N; i++ {
		min_operations.MinOperations(nums, x)
	}
}
