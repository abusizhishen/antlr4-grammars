// Code generated from DCM_2_0_grammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dcm_2_0_grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 42, 523,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 7, 35, 422, 10, 35,
	12, 35, 14, 35, 425, 11, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 7, 37,
	432, 10, 37, 12, 37, 14, 37, 435, 11, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 40, 5, 40, 445, 10, 40, 3, 40, 3, 40, 3, 40, 7, 40,
	450, 10, 40, 12, 40, 14, 40, 453, 11, 40, 5, 40, 455, 10, 40, 3, 41, 5,
	41, 458, 10, 41, 3, 41, 6, 41, 461, 10, 41, 13, 41, 14, 41, 462, 3, 41,
	3, 41, 7, 41, 467, 10, 41, 12, 41, 14, 41, 470, 11, 41, 3, 41, 5, 41, 473,
	10, 41, 3, 41, 5, 41, 476, 10, 41, 3, 41, 3, 41, 6, 41, 480, 10, 41, 13,
	41, 14, 41, 481, 3, 41, 5, 41, 485, 10, 41, 3, 41, 5, 41, 488, 10, 41,
	3, 41, 6, 41, 491, 10, 41, 13, 41, 14, 41, 492, 3, 41, 5, 41, 496, 10,
	41, 3, 42, 3, 42, 3, 43, 3, 43, 5, 43, 502, 10, 43, 3, 43, 6, 43, 505,
	10, 43, 13, 43, 14, 43, 506, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	7, 45, 515, 10, 45, 12, 45, 14, 45, 518, 11, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 516, 2, 46, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28,
	55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 2,
	73, 37, 75, 2, 77, 2, 79, 38, 81, 39, 83, 40, 85, 2, 87, 41, 89, 42, 3,
	2, 10, 6, 2, 48, 48, 50, 59, 93, 93, 95, 95, 5, 2, 67, 92, 97, 97, 99,
	124, 4, 2, 36, 36, 94, 94, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104,
	104, 112, 112, 116, 116, 118, 118, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45,
	47, 47, 5, 2, 11, 11, 14, 15, 34, 34, 5, 2, 35, 35, 44, 44, 48, 48, 2,
	539, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2,
	87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 3, 91, 3, 2, 2, 2, 5, 93, 3, 2, 2, 2,
	7, 118, 3, 2, 2, 2, 9, 128, 3, 2, 2, 2, 11, 139, 3, 2, 2, 2, 13, 143, 3,
	2, 2, 2, 15, 147, 3, 2, 2, 2, 17, 166, 3, 2, 2, 2, 19, 176, 3, 2, 2, 2,
	21, 185, 3, 2, 2, 2, 23, 190, 3, 2, 2, 2, 25, 195, 3, 2, 2, 2, 27, 210,
	3, 2, 2, 2, 29, 220, 3, 2, 2, 2, 31, 234, 3, 2, 2, 2, 33, 251, 3, 2, 2,
	2, 35, 260, 3, 2, 2, 2, 37, 273, 3, 2, 2, 2, 39, 289, 3, 2, 2, 2, 41, 313,
	3, 2, 2, 2, 43, 324, 3, 2, 2, 2, 45, 334, 3, 2, 2, 2, 47, 344, 3, 2, 2,
	2, 49, 354, 3, 2, 2, 2, 51, 363, 3, 2, 2, 2, 53, 375, 3, 2, 2, 2, 55, 379,
	3, 2, 2, 2, 57, 381, 3, 2, 2, 2, 59, 383, 3, 2, 2, 2, 61, 392, 3, 2, 2,
	2, 63, 397, 3, 2, 2, 2, 65, 405, 3, 2, 2, 2, 67, 410, 3, 2, 2, 2, 69, 418,
	3, 2, 2, 2, 71, 426, 3, 2, 2, 2, 73, 428, 3, 2, 2, 2, 75, 438, 3, 2, 2,
	2, 77, 441, 3, 2, 2, 2, 79, 444, 3, 2, 2, 2, 81, 495, 3, 2, 2, 2, 83, 497,
	3, 2, 2, 2, 85, 499, 3, 2, 2, 2, 87, 508, 3, 2, 2, 2, 89, 512, 3, 2, 2,
	2, 91, 92, 7, 12, 2, 2, 92, 4, 3, 2, 2, 2, 93, 94, 7, 77, 2, 2, 94, 95,
	7, 81, 2, 2, 95, 96, 7, 80, 2, 2, 96, 97, 7, 85, 2, 2, 97, 98, 7, 71, 2,
	2, 98, 99, 7, 84, 2, 2, 99, 100, 7, 88, 2, 2, 100, 101, 7, 75, 2, 2, 101,
	102, 7, 71, 2, 2, 102, 103, 7, 84, 2, 2, 103, 104, 7, 87, 2, 2, 104, 105,
	7, 80, 2, 2, 105, 106, 7, 73, 2, 2, 106, 107, 7, 97, 2, 2, 107, 108, 7,
	72, 2, 2, 108, 109, 7, 81, 2, 2, 109, 110, 7, 84, 2, 2, 110, 111, 7, 79,
	2, 2, 111, 112, 7, 67, 2, 2, 112, 113, 7, 86, 2, 2, 113, 114, 7, 34, 2,
	2, 114, 115, 7, 52, 2, 2, 115, 116, 7, 48, 2, 2, 116, 117, 7, 50, 2, 2,
	117, 6, 3, 2, 2, 2, 118, 119, 7, 79, 2, 2, 119, 120, 7, 81, 2, 2, 120,
	121, 7, 70, 2, 2, 121, 122, 7, 87, 2, 2, 122, 123, 7, 78, 2, 2, 123, 124,
	7, 77, 2, 2, 124, 125, 7, 81, 2, 2, 125, 126, 7, 82, 2, 2, 126, 127, 7,
	72, 2, 2, 127, 8, 3, 2, 2, 2, 128, 129, 7, 72, 2, 2, 129, 130, 7, 87, 2,
	2, 130, 131, 7, 80, 2, 2, 131, 132, 7, 77, 2, 2, 132, 133, 7, 86, 2, 2,
	133, 134, 7, 75, 2, 2, 134, 135, 7, 81, 2, 2, 135, 136, 7, 80, 2, 2, 136,
	137, 7, 71, 2, 2, 137, 138, 7, 80, 2, 2, 138, 10, 3, 2, 2, 2, 139, 140,
	7, 71, 2, 2, 140, 141, 7, 80, 2, 2, 141, 142, 7, 70, 2, 2, 142, 12, 3,
	2, 2, 2, 143, 144, 7, 72, 2, 2, 144, 145, 7, 77, 2, 2, 145, 146, 7, 86,
	2, 2, 146, 14, 3, 2, 2, 2, 147, 148, 7, 88, 2, 2, 148, 149, 7, 67, 2, 2,
	149, 150, 7, 84, 2, 2, 150, 151, 7, 75, 2, 2, 151, 152, 7, 67, 2, 2, 152,
	153, 7, 80, 2, 2, 153, 154, 7, 86, 2, 2, 154, 155, 7, 71, 2, 2, 155, 156,
	7, 80, 2, 2, 156, 157, 7, 77, 2, 2, 157, 158, 7, 81, 2, 2, 158, 159, 7,
	70, 2, 2, 159, 160, 7, 75, 2, 2, 160, 161, 7, 71, 2, 2, 161, 162, 7, 84,
	2, 2, 162, 163, 7, 87, 2, 2, 163, 164, 7, 80, 2, 2, 164, 165, 7, 73, 2,
	2, 165, 16, 3, 2, 2, 2, 166, 167, 7, 77, 2, 2, 167, 168, 7, 84, 2, 2, 168,
	169, 7, 75, 2, 2, 169, 170, 7, 86, 2, 2, 170, 171, 7, 71, 2, 2, 171, 172,
	7, 84, 2, 2, 172, 173, 7, 75, 2, 2, 173, 174, 7, 87, 2, 2, 174, 175, 7,
	79, 2, 2, 175, 18, 3, 2, 2, 2, 176, 177, 7, 72, 2, 2, 177, 178, 7, 71,
	2, 2, 178, 179, 7, 85, 2, 2, 179, 180, 7, 86, 2, 2, 180, 181, 7, 89, 2,
	2, 181, 182, 7, 71, 2, 2, 182, 183, 7, 84, 2, 2, 183, 184, 7, 86, 2, 2,
	184, 20, 3, 2, 2, 2, 185, 186, 7, 89, 2, 2, 186, 187, 7, 71, 2, 2, 187,
	188, 7, 84, 2, 2, 188, 189, 7, 86, 2, 2, 189, 22, 3, 2, 2, 2, 190, 191,
	7, 86, 2, 2, 191, 192, 7, 71, 2, 2, 192, 193, 7, 90, 2, 2, 193, 194, 7,
	86, 2, 2, 194, 24, 3, 2, 2, 2, 195, 196, 7, 72, 2, 2, 196, 197, 7, 71,
	2, 2, 197, 198, 7, 85, 2, 2, 198, 199, 7, 86, 2, 2, 199, 200, 7, 89, 2,
	2, 200, 201, 7, 71, 2, 2, 201, 202, 7, 84, 2, 2, 202, 203, 7, 86, 2, 2,
	203, 204, 7, 71, 2, 2, 204, 205, 7, 68, 2, 2, 205, 206, 7, 78, 2, 2, 206,
	207, 7, 81, 2, 2, 207, 208, 7, 69, 2, 2, 208, 209, 7, 77, 2, 2, 209, 26,
	3, 2, 2, 2, 210, 211, 7, 77, 2, 2, 211, 212, 7, 71, 2, 2, 212, 213, 7,
	80, 2, 2, 213, 214, 7, 80, 2, 2, 214, 215, 7, 78, 2, 2, 215, 216, 7, 75,
	2, 2, 216, 217, 7, 80, 2, 2, 217, 218, 7, 75, 2, 2, 218, 219, 7, 71, 2,
	2, 219, 28, 3, 2, 2, 2, 220, 221, 7, 72, 2, 2, 221, 222, 7, 71, 2, 2, 222,
	223, 7, 85, 2, 2, 223, 224, 7, 86, 2, 2, 224, 225, 7, 77, 2, 2, 225, 226,
	7, 71, 2, 2, 226, 227, 7, 80, 2, 2, 227, 228, 7, 80, 2, 2, 228, 229, 7,
	78, 2, 2, 229, 230, 7, 75, 2, 2, 230, 231, 7, 80, 2, 2, 231, 232, 7, 75,
	2, 2, 232, 233, 7, 71, 2, 2, 233, 30, 3, 2, 2, 2, 234, 235, 7, 73, 2, 2,
	235, 236, 7, 84, 2, 2, 236, 237, 7, 87, 2, 2, 237, 238, 7, 82, 2, 2, 238,
	239, 7, 82, 2, 2, 239, 240, 7, 71, 2, 2, 240, 241, 7, 80, 2, 2, 241, 242,
	7, 77, 2, 2, 242, 243, 7, 71, 2, 2, 243, 244, 7, 80, 2, 2, 244, 245, 7,
	80, 2, 2, 245, 246, 7, 78, 2, 2, 246, 247, 7, 75, 2, 2, 247, 248, 7, 80,
	2, 2, 248, 249, 7, 75, 2, 2, 249, 250, 7, 71, 2, 2, 250, 32, 3, 2, 2, 2,
	251, 252, 7, 77, 2, 2, 252, 253, 7, 71, 2, 2, 253, 254, 7, 80, 2, 2, 254,
	255, 7, 80, 2, 2, 255, 256, 7, 72, 2, 2, 256, 257, 7, 71, 2, 2, 257, 258,
	7, 78, 2, 2, 258, 259, 7, 70, 2, 2, 259, 34, 3, 2, 2, 2, 260, 261, 7, 72,
	2, 2, 261, 262, 7, 71, 2, 2, 262, 263, 7, 85, 2, 2, 263, 264, 7, 86, 2,
	2, 264, 265, 7, 77, 2, 2, 265, 266, 7, 71, 2, 2, 266, 267, 7, 80, 2, 2,
	267, 268, 7, 80, 2, 2, 268, 269, 7, 72, 2, 2, 269, 270, 7, 71, 2, 2, 270,
	271, 7, 78, 2, 2, 271, 272, 7, 70, 2, 2, 272, 36, 3, 2, 2, 2, 273, 274,
	7, 73, 2, 2, 274, 275, 7, 84, 2, 2, 275, 276, 7, 87, 2, 2, 276, 277, 7,
	82, 2, 2, 277, 278, 7, 82, 2, 2, 278, 279, 7, 71, 2, 2, 279, 280, 7, 80,
	2, 2, 280, 281, 7, 77, 2, 2, 281, 282, 7, 71, 2, 2, 282, 283, 7, 80, 2,
	2, 283, 284, 7, 80, 2, 2, 284, 285, 7, 72, 2, 2, 285, 286, 7, 71, 2, 2,
	286, 287, 7, 78, 2, 2, 287, 288, 7, 70, 2, 2, 288, 38, 3, 2, 2, 2, 289,
	290, 7, 85, 2, 2, 290, 291, 7, 86, 2, 2, 291, 292, 7, 87, 2, 2, 292, 293,
	7, 71, 2, 2, 293, 294, 7, 86, 2, 2, 294, 295, 7, 92, 2, 2, 295, 296, 7,
	85, 2, 2, 296, 297, 7, 86, 2, 2, 297, 298, 7, 71, 2, 2, 298, 299, 7, 78,
	2, 2, 299, 300, 7, 78, 2, 2, 300, 301, 7, 71, 2, 2, 301, 302, 7, 80, 2,
	2, 302, 303, 7, 88, 2, 2, 303, 304, 7, 71, 2, 2, 304, 305, 7, 84, 2, 2,
	305, 306, 7, 86, 2, 2, 306, 307, 7, 71, 2, 2, 307, 308, 7, 75, 2, 2, 308,
	309, 7, 78, 2, 2, 309, 310, 7, 87, 2, 2, 310, 311, 7, 80, 2, 2, 311, 312,
	7, 73, 2, 2, 312, 40, 3, 2, 2, 2, 313, 314, 7, 86, 2, 2, 314, 315, 7, 71,
	2, 2, 315, 316, 7, 90, 2, 2, 316, 317, 7, 86, 2, 2, 317, 318, 7, 85, 2,
	2, 318, 319, 7, 86, 2, 2, 319, 320, 7, 84, 2, 2, 320, 321, 7, 75, 2, 2,
	321, 322, 7, 80, 2, 2, 322, 323, 7, 73, 2, 2, 323, 42, 3, 2, 2, 2, 324,
	325, 7, 71, 2, 2, 325, 326, 7, 75, 2, 2, 326, 327, 7, 80, 2, 2, 327, 328,
	7, 74, 2, 2, 328, 329, 7, 71, 2, 2, 329, 330, 7, 75, 2, 2, 330, 331, 7,
	86, 2, 2, 331, 332, 7, 97, 2, 2, 332, 333, 7, 90, 2, 2, 333, 44, 3, 2,
	2, 2, 334, 335, 7, 71, 2, 2, 335, 336, 7, 75, 2, 2, 336, 337, 7, 80, 2,
	2, 337, 338, 7, 74, 2, 2, 338, 339, 7, 71, 2, 2, 339, 340, 7, 75, 2, 2,
	340, 341, 7, 86, 2, 2, 341, 342, 7, 97, 2, 2, 342, 343, 7, 91, 2, 2, 343,
	46, 3, 2, 2, 2, 344, 345, 7, 71, 2, 2, 345, 346, 7, 75, 2, 2, 346, 347,
	7, 80, 2, 2, 347, 348, 7, 74, 2, 2, 348, 349, 7, 71, 2, 2, 349, 350, 7,
	75, 2, 2, 350, 351, 7, 86, 2, 2, 351, 352, 7, 97, 2, 2, 352, 353, 7, 89,
	2, 2, 353, 48, 3, 2, 2, 2, 354, 355, 7, 78, 2, 2, 355, 356, 7, 67, 2, 2,
	356, 357, 7, 80, 2, 2, 357, 358, 7, 73, 2, 2, 358, 359, 7, 80, 2, 2, 359,
	360, 7, 67, 2, 2, 360, 361, 7, 79, 2, 2, 361, 362, 7, 71, 2, 2, 362, 50,
	3, 2, 2, 2, 363, 364, 7, 70, 2, 2, 364, 365, 7, 75, 2, 2, 365, 366, 7,
	85, 2, 2, 366, 367, 7, 82, 2, 2, 367, 368, 7, 78, 2, 2, 368, 369, 7, 67,
	2, 2, 369, 370, 7, 91, 2, 2, 370, 371, 7, 80, 2, 2, 371, 372, 7, 67, 2,
	2, 372, 373, 7, 79, 2, 2, 373, 374, 7, 71, 2, 2, 374, 52, 3, 2, 2, 2, 375,
	376, 7, 88, 2, 2, 376, 377, 7, 67, 2, 2, 377, 378, 7, 84, 2, 2, 378, 54,
	3, 2, 2, 2, 379, 380, 7, 46, 2, 2, 380, 56, 3, 2, 2, 2, 381, 382, 7, 63,
	2, 2, 382, 58, 3, 2, 2, 2, 383, 384, 7, 72, 2, 2, 384, 385, 7, 87, 2, 2,
	385, 386, 7, 80, 2, 2, 386, 387, 7, 77, 2, 2, 387, 388, 7, 86, 2, 2, 388,
	389, 7, 75, 2, 2, 389, 390, 7, 81, 2, 2, 390, 391, 7, 80, 2, 2, 391, 60,
	3, 2, 2, 2, 392, 393, 7, 85, 2, 2, 393, 394, 7, 86, 2, 2, 394, 395, 7,
	49, 2, 2, 395, 396, 7, 90, 2, 2, 396, 62, 3, 2, 2, 2, 397, 398, 7, 85,
	2, 2, 398, 399, 7, 86, 2, 2, 399, 400, 7, 97, 2, 2, 400, 401, 7, 86, 2,
	2, 401, 402, 7, 90, 2, 2, 402, 403, 7, 49, 2, 2, 403, 404, 7, 90, 2, 2,
	404, 64, 3, 2, 2, 2, 405, 406, 7, 85, 2, 2, 406, 407, 7, 86, 2, 2, 407,
	408, 7, 49, 2, 2, 408, 409, 7, 91, 2, 2, 409, 66, 3, 2, 2, 2, 410, 411,
	7, 85, 2, 2, 411, 412, 7, 86, 2, 2, 412, 413, 7, 97, 2, 2, 413, 414, 7,
	86, 2, 2, 414, 415, 7, 90, 2, 2, 415, 416, 7, 49, 2, 2, 416, 417, 7, 91,
	2, 2, 417, 68, 3, 2, 2, 2, 418, 423, 5, 71, 36, 2, 419, 422, 5, 71, 36,
	2, 420, 422, 9, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 420, 3, 2, 2, 2, 422,
	425, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 70, 3,
	2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 427, 9, 3, 2, 2, 427, 72, 3, 2, 2,
	2, 428, 433, 7, 36, 2, 2, 429, 432, 5, 75, 38, 2, 430, 432, 10, 4, 2, 2,
	431, 429, 3, 2, 2, 2, 431, 430, 3, 2, 2, 2, 432, 435, 3, 2, 2, 2, 433,
	431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 436, 3, 2, 2, 2, 435, 433,
	3, 2, 2, 2, 436, 437, 7, 36, 2, 2, 437, 74, 3, 2, 2, 2, 438, 439, 7, 94,
	2, 2, 439, 440, 9, 5, 2, 2, 440, 76, 3, 2, 2, 2, 441, 442, 7, 36, 2, 2,
	442, 78, 3, 2, 2, 2, 443, 445, 5, 83, 42, 2, 444, 443, 3, 2, 2, 2, 444,
	445, 3, 2, 2, 2, 445, 454, 3, 2, 2, 2, 446, 455, 7, 50, 2, 2, 447, 451,
	4, 51, 59, 2, 448, 450, 4, 50, 59, 2, 449, 448, 3, 2, 2, 2, 450, 453, 3,
	2, 2, 2, 451, 449, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 455, 3, 2, 2,
	2, 453, 451, 3, 2, 2, 2, 454, 446, 3, 2, 2, 2, 454, 447, 3, 2, 2, 2, 455,
	80, 3, 2, 2, 2, 456, 458, 5, 83, 42, 2, 457, 456, 3, 2, 2, 2, 457, 458,
	3, 2, 2, 2, 458, 460, 3, 2, 2, 2, 459, 461, 4, 50, 59, 2, 460, 459, 3,
	2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 460, 3, 2, 2, 2, 462, 463, 3, 2, 2,
	2, 463, 464, 3, 2, 2, 2, 464, 468, 7, 48, 2, 2, 465, 467, 4, 50, 59, 2,
	466, 465, 3, 2, 2, 2, 467, 470, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 468,
	469, 3, 2, 2, 2, 469, 472, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 471, 473,
	5, 85, 43, 2, 472, 471, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 496, 3,
	2, 2, 2, 474, 476, 5, 83, 42, 2, 475, 474, 3, 2, 2, 2, 475, 476, 3, 2,
	2, 2, 476, 477, 3, 2, 2, 2, 477, 479, 7, 48, 2, 2, 478, 480, 4, 50, 59,
	2, 479, 478, 3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 481,
	482, 3, 2, 2, 2, 482, 484, 3, 2, 2, 2, 483, 485, 5, 85, 43, 2, 484, 483,
	3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 496, 3, 2, 2, 2, 486, 488, 5, 83,
	42, 2, 487, 486, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 490, 3, 2, 2, 2,
	489, 491, 4, 50, 59, 2, 490, 489, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492,
	490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 496,
	5, 85, 43, 2, 495, 457, 3, 2, 2, 2, 495, 475, 3, 2, 2, 2, 495, 487, 3,
	2, 2, 2, 496, 82, 3, 2, 2, 2, 497, 498, 7, 47, 2, 2, 498, 84, 3, 2, 2,
	2, 499, 501, 9, 6, 2, 2, 500, 502, 9, 7, 2, 2, 501, 500, 3, 2, 2, 2, 501,
	502, 3, 2, 2, 2, 502, 504, 3, 2, 2, 2, 503, 505, 4, 50, 59, 2, 504, 503,
	3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 504, 3, 2, 2, 2, 506, 507, 3, 2,
	2, 2, 507, 86, 3, 2, 2, 2, 508, 509, 9, 8, 2, 2, 509, 510, 3, 2, 2, 2,
	510, 511, 8, 44, 2, 2, 511, 88, 3, 2, 2, 2, 512, 516, 9, 9, 2, 2, 513,
	515, 11, 2, 2, 2, 514, 513, 3, 2, 2, 2, 515, 518, 3, 2, 2, 2, 516, 517,
	3, 2, 2, 2, 516, 514, 3, 2, 2, 2, 517, 519, 3, 2, 2, 2, 518, 516, 3, 2,
	2, 2, 519, 520, 7, 12, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 8, 45, 2,
	2, 522, 90, 3, 2, 2, 2, 23, 2, 421, 423, 431, 433, 444, 451, 454, 457,
	462, 468, 472, 475, 481, 484, 487, 492, 495, 501, 506, 516, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\n'", "'KONSERVIERUNG_FORMAT 2.0'", "'MODULKOPF'", "'FUNKTIONEN'",
	"'END'", "'FKT'", "'VARIANTENKODIERUNG'", "'KRITERIUM'", "'FESTWERT'",
	"'WERT'", "'TEXT'", "'FESTWERTEBLOCK'", "'KENNLINIE'", "'FESTKENNLINIE'",
	"'GRUPPENKENNLINIE'", "'KENNFELD'", "'FESTKENNFELD'", "'GRUPPENKENNFELD'",
	"'STUETZSTELLENVERTEILUNG'", "'TEXTSTRING'", "'EINHEIT_X'", "'EINHEIT_Y'",
	"'EINHEIT_W'", "'LANGNAME'", "'DISPLAYNAME'", "'VAR'", "','", "'='", "'FUNKTION'",
	"'ST/X'", "'ST_TX/X'", "'ST/Y'", "'ST_TX/Y'", "", "", "", "", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NAME",
	"TEXT", "INT", "FLOAT", "MINUS", "WS", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"NAME", "LETTER", "TEXT", "EscapeSequence", "QUOTE", "INT", "FLOAT", "MINUS",
	"Exponent", "WS", "COMMENT",
}

type DCM_2_0_grammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDCM_2_0_grammarLexer(input antlr.CharStream) *DCM_2_0_grammarLexer {

	l := new(DCM_2_0_grammarLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "DCM_2_0_grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DCM_2_0_grammarLexer tokens.
const (
	DCM_2_0_grammarLexerT__0    = 1
	DCM_2_0_grammarLexerT__1    = 2
	DCM_2_0_grammarLexerT__2    = 3
	DCM_2_0_grammarLexerT__3    = 4
	DCM_2_0_grammarLexerT__4    = 5
	DCM_2_0_grammarLexerT__5    = 6
	DCM_2_0_grammarLexerT__6    = 7
	DCM_2_0_grammarLexerT__7    = 8
	DCM_2_0_grammarLexerT__8    = 9
	DCM_2_0_grammarLexerT__9    = 10
	DCM_2_0_grammarLexerT__10   = 11
	DCM_2_0_grammarLexerT__11   = 12
	DCM_2_0_grammarLexerT__12   = 13
	DCM_2_0_grammarLexerT__13   = 14
	DCM_2_0_grammarLexerT__14   = 15
	DCM_2_0_grammarLexerT__15   = 16
	DCM_2_0_grammarLexerT__16   = 17
	DCM_2_0_grammarLexerT__17   = 18
	DCM_2_0_grammarLexerT__18   = 19
	DCM_2_0_grammarLexerT__19   = 20
	DCM_2_0_grammarLexerT__20   = 21
	DCM_2_0_grammarLexerT__21   = 22
	DCM_2_0_grammarLexerT__22   = 23
	DCM_2_0_grammarLexerT__23   = 24
	DCM_2_0_grammarLexerT__24   = 25
	DCM_2_0_grammarLexerT__25   = 26
	DCM_2_0_grammarLexerT__26   = 27
	DCM_2_0_grammarLexerT__27   = 28
	DCM_2_0_grammarLexerT__28   = 29
	DCM_2_0_grammarLexerT__29   = 30
	DCM_2_0_grammarLexerT__30   = 31
	DCM_2_0_grammarLexerT__31   = 32
	DCM_2_0_grammarLexerT__32   = 33
	DCM_2_0_grammarLexerNAME    = 34
	DCM_2_0_grammarLexerTEXT    = 35
	DCM_2_0_grammarLexerINT     = 36
	DCM_2_0_grammarLexerFLOAT   = 37
	DCM_2_0_grammarLexerMINUS   = 38
	DCM_2_0_grammarLexerWS      = 39
	DCM_2_0_grammarLexerCOMMENT = 40
)
