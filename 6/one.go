package main



func main(){

input := `jghhttcmttdwwfjjjpqjqllwvwffswwqmwmddvndvdrrcwrccfncfcgcjjpjrjzjggczgglhllbzlztlzzswwbqqrvrvqrqcrclrccrbrpbbzhbbzvbbwrwswqwqqgmmgqmqmdqdhdmmfnndpdvvdvdbbmgmqgqgqmqrrsbsrrzcrzrcrvrddpqdpdwwsvshvshvshszhshddjllbjljjwhjwhwppbssdccvzzmvvwrrrhmmbdbggqhgqglglgrrdbdjddjvddzpznzmnnpgnppnznzvzbzvvdtdwdttnzzdpzzqjzjzllhglghgttcbtcbtbjbssqcqtcqtqvttvpvqvggbsszjsjqjllfgfqqhllvsvmmjhhbdbzzrwzzggdfgdgcddgmmjnjhhhgvgwghgfffclffdmmmlclzclzzqgqjqzqwzqzbbbljlplnlrlqlqlhlvlmvvlwltwtqwwznndpdttjjrllpptnngjnnzppjhpjhpjpjmjpmjpptzppjpsprpbpmmshmsmqqnhnjhnjnccdnnqhhgfgmgnmnvvtwvttcbtbnbcnnhjjnpphjpjppshsppmrmrbrlrzlzjlzzvssvfffsjfsjshshppchppprpcclhlmmgwwsddgbdggcqqhsqqbdbzzdffmpmvmpptppjrprrgvvgrrztrtllnblbffpgprpqqwcwjwcjcbbjgbbsgsbgbgwbwrrbcrcjcnnvlnlmnnwcncbcrrgcrggtqqgbbqsqhsqslswsdswscwssbqbppczzpmmnjnfjfpjpqpddpbbgrrnqrnqqnwqnntsnsdnddvhhgrgmmgzmgmwgmwwmfwwnqnjjgqghqhhpzpmpbmpbbqvvjjlvvddcnnglnnhpnhhjghgrgjrjqqgtqggmzzbdbssbpssrjsrsbbpzzsgzzpcpwwfgwfgfccslsmlltlnttpvvcbvbqqqqvlvplvvmnntstffvhvvwfftltjjtftbftfqqnbbmnmqqwbbldbldbddtldtldldbllsggdqgdgfgtggglmmgfgcfgccgttgdgdvgglhhvbbnzbzvbzzvtvntvntvvwgvgcgrrmbmnmttzlzgllpmlpmmscsncnnnfrrvqvmvsvsrvvbvpbbvsvssdnnmvnnnggpmpbbqsqwssjwwtnwwhbwwmcwwjfwfhfvvsnsqqwmmrfmrmcrchhpqprpzpmpzzqqmqbqjjbffjnnhphdpdjjbcjcsjcjcdcrddscsgsqgsqggcscrssscvcncfccptcccjgjqqlppjbpjpgpzpgpqqzwqwcwnwccmnnnwrnrttjqtjtllbslsrlrggjrrjwrwmrrcbbnddzcdzcchqqrhhcrcrlrmrcchcpcmcrrlqltlhttdhhjrrzllwjjfvvzccmwwtlwwhpwwssfdsfdsfsddwbbfsfqfwqffhqfqvqwwhfhrrwhrwhrwrzrwrzwwbsblslzzrwzwdwnwznnbsnsbbgffjppvqppbmmdffgghrrchrhfhthqthqhvvrcvrrfvfbvffmmbjmjbbwhhvqvbvjjfzjzvvwnvncvvdssfffjfhjhbbdhdlllmrlmrllbtltbbhnnczzwjzwjwllvtvmvpvvmmfnmnwmwjwrrqtqlltwllwrllsvvdrvrnvrnnsjnjfnnmpnpgnglnggbnbmmcmddhrdrjdjtthssdbssjtjvttqwqppwbpblljnlnljlfjffvddtzzqfzqqjttqhhhvppghpghpghghjggsrrczzqvvbjbbbglldcdrdqdpqpwpffnjnhnthhtlhhrdrcrwrvrlvrlvrvprpccmffqzffdbdhhwjjgmgzzqhqnqsstltjjzszpsppdpnddrnrmrllhtwnlhgzgvsfjfmgfcnnzqhbfztnzhnctmjjhvzrhjcptzqqtstlmrgnbcpnctjgswhfcmtzgsndfzdlqsrlcjrssmjndgrzgvtgfjwtcqwnzmrgtgngcwcwzvttqcdwqspzldfmzgmfclwgqvvvqhhhswpdjlbpjpppdljbdjrbblbrtftdsmvmfpftdlvhdphzbcwwwjvmsjczbtblwbtszmbcmpbdqnwcgcltbdzbsvtjhlqgrsgqzztqrvmswtzgvsjslgvjcvhqftdcmwqwhgwrmrpfdqvfqhczlztrhjqlppchgspfjwzvfsncdhgnlnnshwsbvqmqwtldtmshhqpqbdzqlwfvvbbzwqvlvqdwcpbtrsrhmrjnzmqjttthhbbdjwtzhwjcvhzrtgfwltqhmzwtqbgjjpphlfwfhgctmwpqrmngfrfmbmcqpqcrsmzhjzdzbppbfwpgdtdjvjdvbwfcbddjpjbgtrjtddlpvnppcwhbbgpqgmcrwmfsvqspdvctvljbhcgnllzdpjmjdqwflrfqrsfqnhcsbtszjfmqvjwmcsghwcssbnzzpsggbmpdqrlctrcrmbzlnnqcfmrpfwsnjcggtqncwddsjjzwvlnfpwjlrgwrrvfhgtjwqlblnsnrddjqqtwtzrvfqfqcjhbwtlmfsfgpzgtdddtlvqlqgjwpprpzvdhvlfjbqqflbnvgsgblfpcqprlqrhrrddjjftbmfgghrrrhmttmnfzgmrzwcscnsmdnnddsbdjswhwmjfjnfvjbtcqjlflzjsqqhldcdsbjttmvmcdbwrfwdlllbbjhbmdhnfgwmmbpbsrmqptppqzwtncnwwsjjrchgpvdcsspfqpczmsqfrgvfgsnhfmplrnhzddlbvnthltpwbbzdwwnmhmllwfphhfpnghgdzlzgpbvsphbhvfmcvqpqvdsjjbnbsdvccbmgwdbgsnhpgfshzzznjdbngsrqpwhqjqdfsnmngzznwgvqrtjmzcmbbqjgjcdjdqbmdrgvqflsstqhgsgpfnmzvrdgztlzlhvdjdcslwbgldwjwftvczvtbpdtdqtqshqrbpvpgbfbtnnlmzrhdtsjdlllbtwqgcfphssqmszmbllnbpwvhfdllwtcbqccmbtscmsvppjjrcpswgtzgvhblvbmcddhhgqghbtwzffscsvzgwjdfldccbcfptpmmfbwqzlqhdrcdvhpnwvddqcrwwlmtgvrcvlbvtblhhmhdrvvnpmrdqsgjdrfprfcmtfsbrmsglfjcnjwvpjqlptrbmrqcrdfccmfvhqzrmwqbbmtmjrbnmvzdbwcfpwjzrjrhzrncpwptswhnrgsdpdfjqcjhnvvftgzdccpsgdqcqzfvbtftrzljqmjbgmmlrdlvpwqddmdhzsslzlnvfrblfzfpwtvhpwhmqmbzvchndbzfswtmvcprhlssmncfdqcpzjczptptgdzpjvrqtrlcgbhlqqwvnsrrvtllldbhztgnlmjbmqmgcpcbwtthnhwghnqgvqtnnltbzslmmbjbhqtrmqbgccphjsgwbtstvjnpjjqsdwdsgcrjgznntsgvlrgzmbnzdmbghrmtscppgmztfbsqczzvhsmjmmclhrcbgnvjbgffgbrhrdwccplpdfzcljgfntjlzmsqrwrmdqbclffbfpscddcfsbpsnnphjvjwsfvbprsgpcnbrblvvqfvngbhgmmglzzfmmpnnrphwrvqnmffwpsmmrlhmnsdvbdgjlrjmsdzjltgmjfplrrfmgbhzltzzfpwtqcqwbgsgwfbmntzsvtqqtnrhbtbnshfmwwzbvtsqmtsfssrcvgngvvhjlcnfsvltpsdjspgmmrqttcsltjzqglpspmgrnbrtnbtnjzprqmtfhgczjrlqdhsjqjbhpnwhrffmhvqfmlzcpfflptgqwthfzvspbjjdwmmbnttbztzpnjmlstgmgqbptdncqgbdbdnqwslwhfrdfvmbbqlzhdptpfrhnvcchpddngslzzrhsrwclpccbqhcscbzcpdtwmppvrpjnnjfgrswndtzprjnsvvdwwhhbcsglnwwptptdzgsmbwppdrhwpqhzlgfcsqtfzvqdvcsbtbqvtfvwlcdrwttgmwhbjlqphclqfzmlb`
runeArray := []rune(input)
// left:=0;

result := 0;
for c:=0;c<len(runeArray);c++{
	if(c > 2){
		if c == 3{
			print(string(runeArray[c]))
		}

		check := string(runeArray[c]) + string(runeArray[c-1]) + string(runeArray[c-2]) +string(runeArray[c-3])

		m := make(map[rune]bool)
		unique := true
    	for _, i := range check {
        _, ok := m[i]
        if ok {
            unique = false
        }
        m[i] = true
    }
		if unique{
			result = c + 1
			break
		}

	}
}

print(result)


}