package porter

/* This is the Porter stemming algorithm, coded up as thread-safe ANSI C
   by the author.

   It may be be regarded as cononical, in that it follows the algorithm
   presented in

   Porter, 1980, An algorithm for suffix stripping, Program, Vol. 14,
   no. 3, pp 130-137,

   only differing from it at the points maked --DEPARTURE-- below.

   See also http://www.tartarus.org/~martin/PorterStemmer

   The algorithm as described in the paper could be exactly replicated
   by adjusting the points of DEPARTURE, but this is barely necessary,
   because (a) the points of DEPARTURE are definitely improvements, and
   (b) no encoding of the Porter stemmer I have seen is anything like
   as exact as this version, even with the points of DEPARTURE!

   You can compile it on Unix with 'gcc -O3 -o stem stem.c' after which
   'stem' takes a list of inputs and sends the stemmed equivalent to
   stdout.

   The algorithm as encoded here is particularly fast.

   Release 2 (the more old-fashioned, non-thread-safe version may be
   regarded as release 1.)
*/


import (
	"bytes"
	"strings"
	"fmt"
)


var (
	_BLANK = []byte("")
	_ABLE = []byte("able")
	_AL  = []byte("al")
	_ALISM = []byte("alism")
	_ALITI = []byte("aliti")
	_ALIZE = []byte("alize")
	_ALLI = []byte("alli")
	_ANCE = []byte("ance")
	_ANCI = []byte("anci")
	_ANT = []byte("ant")
	_AT  = []byte("at")
	_ATE = []byte("ate")
	_ATION = []byte("ation")
	_ATIONAL = []byte("ational")
	_ATIVE = []byte("ative")
	_ATOR = []byte("ator")
	_BILITI = []byte("biliti")
	_BL  = []byte("bl")
	_BLE = []byte("ble")
	_BLI = []byte("bli")
	_E   = []byte("e")
	_ED   = []byte("ed")
	_EED  = []byte("eed")
	_ELI = []byte("eli")
	_EMENT = []byte("ement")
	_ENCE = []byte("ence")
	_ENCI = []byte("enci")
	_ENT = []byte("ent")
	_ENTLI = []byte("entli")
	_ER = []byte("er")
	_FUL= []byte("ful")
	_FULNESS = []byte("fulness")
	_I    = []byte("i")
	_IBLE = []byte("ible")
	_IC = []byte("ic")
	_ICAL = []byte("ical")
	_ICATE = []byte("icate")
	_ICITI = []byte("iciti")
	_IES  = []byte("ies")
	_ING  = []byte("ing")
	_ION = []byte("ion")
	_ISM = []byte("ism")
	_ITI = []byte("iti")
	_IVE = []byte("ive")
	_IVENESS = []byte("iveness")
	_IVITI = []byte("iviti")
	_IZ = []byte("iz")
	_IZATION = []byte("ization")
	_IZE = []byte("ize")
	_IZER = []byte("izer")
	_LOG = []byte("log")
	_LOGI = []byte("logi")
	_MENT = []byte("ment")
	_NESS = []byte("ness")
	_OU = []byte("ou")
	_OUS = []byte("ous")
	_OUSLI = []byte("ousli")
	_OUSNESS = []byte("ousness")
	_SSES = []byte("sses")
	_TION = []byte("tion")
	_TIONAL = []byte("tional")
	_Y    = []byte("y")
)


type stemmer struct {
  b   []byte
	j   int
	k   int
}


///* The main part of the stemming algorithm starts here.
//*/
//
//
//
///* Member b is a buffer holding a word to be stemmed. The letters are in
//   b[0], b[1] ... ending at b[z->k]. Member k is readjusted downwards as
//   the stemming progresses. Zero termination is not in fact used in the
//   algorithm.
//
//   Note that only lower case sequences are stemmed. Forcing to lower case
//   should be done before stem(...) is called.
//
//
//   Typical usage is:
//
//       struct stemmer * z = create_stemmer();
//       char b[] = "pencils";
//       int res = stem(z, b, 6);
//           /- stem the 7 characters of b[0] to b[6]. The result, res,
//              will be 5 (the 's' is removed). -/
//       free_stemmer(z);
//*/

/*
 * Check wheter the letter at position i is a consonant
 */
func(z *stemmer) consonant (pos int) (bool) {
  if (len(z.b) <= pos) {
    return false
  }
  switch (z.b[pos]) {
    case 'a': fallthrough
    case 'e': fallthrough
    case 'i': fallthrough
    case 'o': fallthrough
    case 'u':
      return false
    case 'y':
      if pos == 0 {
        return true
      } else {
        return z.vowel(pos-1)
      }
  }
	return true
}

func (z *stemmer) vowel (pos int) bool {
  return !z.consonant(pos)
}


/* m(z) measures the number of consonant sequences between 0 and j. if c is
   a consonant sequence and v a vowel sequence, and <..> indicates arbitrary
   presence,

      <c><v>       gives 0
      <c>vc<v>     gives 1
      <c>vcvc<v>   gives 2
      <c>vcvcvc<v> gives 3
      ....
*/


func (z *stemmer) m ()(int) {
	var n,i int 

	for {
		if i>z.j {
			return n
		}
		if !z.consonant(i) {
			break
		}
		i++
	}
	i++
	for {
		for {
			if ( i>z.j) {
				return n
			}
			if z.consonant(i) {
				break
			}
			i++
		}
		i++
		n++
		for {
			if (i>z.j) {
				return n
			}
			if !z.consonant(i) {
				break
			}
			i++
		}
		i++
	}
	return n
}

/* vowelinstem(z) is TRUE <=> 0,...j contains a vowel */


func (z *stemmer) vowelinstem () (bool) {
	for i:=0; i<=z.j; i++ {
		if !z.consonant(i) {
			return true
		}
	}
	return false
}

/* doublec(z, j) is TRUE <=> j,(j-1) contain a double consonant. */


func (z *stemmer) doublec(j int)(bool){
	if 1 > j {
		return false
	}
	if z.b[j] != z.b[j-1] {
		return false
	}
	return z.consonant(j)

}

func log (msg string, args ...interface{}) {
  fmt.Printf(msg, args...)
}
/* cvc(z, i) is TRUE <=> i-2,i-1,i has the form consonant - vowel - consonant
   and also if the second c is not w,x or y. this is used when trying to
   restore an e at the end of a short word. e.g.

      cav(e), lov(e), hop(e), crim(e), but
      snow, box, tray.

*/


func (z *stemmer)cvc(i int) (bool){
	if	2>i || !z.consonant(i) || z.consonant(i-1) || !z.consonant(i-2) {
			return false
	}
	switch z.b[i] {
		case 'w': fallthrough
		case 'x': fallthrough
		case 'y': 
			return false
	}
	return true
}


/* ends(z, s) is TRUE <=> 0,...k ends with the string s. */
func (z *stemmer) ends(s []byte)(bool) {
	length := len(s)
	//fmt.Printf("%d %d\n", len(z.b), z.k)
	if length > z.k {
		return false
	}
	if !bytes.HasSuffix(z.b[:z.k+1], s) {
		return false
	}
	z.j = z.k-length
	return true
}

/* setto(z, s) sets (j+1),...k to the characters in the string s, readjusting
   k. */


func (z *stemmer) setto (s []byte) {
	j := z.j

	for _, b := range(s) {
		z.b[j+1] = b
		j++
	}
	z.k = j
}

/* r(z, s) is used further down. */

func (z *stemmer) r(s []byte) {
	if 0 < z.m() {
		z.setto(s)
	}
}

/* step1ab(z) gets rid of plurals and -ed or -ing. e.g.

       caresses  ->  caress
       ponies    ->  poni
       ties      ->  ti
       caress    ->  caress
       cats      ->  cat

       feed      ->  feed
       agreed    ->  agree
       disabled  ->  disable

       matting   ->  mat
       mating    ->  mate
       meeting   ->  meet
       milling   ->  mill
       messing   ->  mess

       meetings  ->  meet

*/


func (z *stemmer) step1ab () {
	if 's' == z.b[z.k] {
		switch {
			case z.ends(_SSES):
				z.k -= 2
			case z.ends(_IES):
				z.setto(_I)
			default:
				if 's' != z.b[z.k-1] {
					z.k--
				}
		}
	}
	if z.ends(_EED) {
		if 0 < z.m() {
			z.k--
		}
	} else if (z.ends(_ED) || z.ends(_ING)) && z.vowelinstem() {
		z.k = z.j
		switch {
			case z.ends(_AT):
				z.setto(_ATE)
			case z.ends(_BL):
				z.setto(_BLE)
			case z.ends(_IZ):
				z.setto(_IZE)
			case z.doublec(z.k):
				z.k--
				switch z.b[z.k]{
					case 'l': fallthrough
					case 's': fallthrough
					case 'z':
						z.k++
				}
			default:
				if 1 == z.m() && z.cvc(z.k) {
					z.setto(_E)
				}
		}
	}
}

/* step1c(z) turns terminal y to i when there is another vowel in the stem. */

func (z *stemmer) step1c() {
	if z.ends(_Y) && z.vowelinstem() {
		z.b[z.k] = 'i'
	}
}


/* step2(z) maps double suffices to single ones. so -ization ( = -ize plus
   -ation) maps to -ize etc. note that the string before the suffix must give
   m(z) > 0. */

func (z *stemmer) step2() {
	switch z.b[z.k-1] {
		case'a':z.step2_a()
		case'c':z.step2_c()
		case'e':z.step2_e()
		case'l':z.step2_l()
		case'o':z.step2_o()
		case's':z.step2_s()
		case't':z.step2_t()
		case'g':z.step2_g()
	}
}
func (z *stemmer) step2_a(){
	switch {
		case z.ends(_ATIONAL): z.r(_ATE)
		case z.ends(_TIONAL):  z.r(_TION)
	}
}
func (z *stemmer) step2_c(){
	switch {
		case z.ends(_ENCI): z.r(_ENCE)
		case z.ends(_ANCI): z.r(_ANCE)
	}
}
func (z *stemmer) step2_e(){
	if z.ends(_IZER) {
		z.r(_IZE)
	}
}
func (z *stemmer) step2_l(){
	switch {
		case z.ends(_BLI):   z.r(_BLE)
		case z.ends(_ALLI):  z.r(_AL)
		case z.ends(_ENTLI): z.r(_ENT)
		case z.ends(_ELI):   z.r(_E)
		case z.ends(_OUSLI): z.r(_OUS)
	}
}
func (z *stemmer) step2_o(){
	switch {
		case z.ends(_IZATION): z.r(_IZE)
		case z.ends(_ATION):   z.r(_ATE)
		case z.ends(_ATOR):    z.r(_ATE)
	}
}

func (z *stemmer) step2_s(){
	switch {
		case z.ends(_ALISM):   z.r(_AL)
		case z.ends(_IVENESS): z.r(_IVE)
		case z.ends(_FULNESS): z.r(_FUL)
		case z.ends(_OUSNESS): z.r(_OUS)
	}
}
func (z *stemmer) step2_t(){
	switch {
		case z.ends(_ALITI):  z.r(_AL)
		case z.ends(_IVITI):  z.r(_IVE)
		case z.ends(_BILITI): z.r(_BLE)
	}
}
func (z *stemmer) step2_g(){
	if z.ends(_LOGI) {
		z.r(_LOG)
	}
}

/* step3(z) deals with -ic-, -full, -ness etc. similar strategy to step2. */


func (z *stemmer) step3 () {
	switch z.b[z.k] {
		case'e':z.step3_e()
		case'i':z.step3_i()
		case'l':z.step3_l()
		case's':z.step3_s()
	}
}

func (z *stemmer) step3_e(){
	switch {
		case z.ends(_ICATE): z.r(_IC)
		case z.ends(_ATIVE): z.r(_BLANK)
		case z.ends(_ALIZE): z.r(_AL)
	}
}
func (z *stemmer) step3_i(){
	if z.ends(_ICITI) {
		z.r(_IC)
	}
}
func (z *stemmer) step3_l(){
	switch {
		case z.ends(_ICAL): z.r(_IC)
		case z.ends(_FUL): z.r(_BLANK)
	}
}
func (z *stemmer) step3_s(){
	if z.ends(_NESS) {
		z.r(_BLANK)
	}
}
/* step4(z) takes off -ant, -ence etc., in context <c>vcvc<v>. */


func (z *stemmer) step4 () {
	switch z.b[z.k-1]{
		case 'a':z.step4_a()
		case 'c':z.step4_c()
		case 'e':z.step4_e()
		case 'i':z.step4_i()
		case 'l':z.step4_l()
		case 'n':z.step4_n()
		case 'o':z.step4_o()
		case 's':z.step4_s()
		case 't':z.step4_t()
		case 'u':z.step4_u()
		case 'v':z.step4_v()
		case 'z':z.step4_z()
	}
}
func (z *stemmer) step4_update() {
	if 1 < z.m() {
		z.k = z.j
	}
}
func (z *stemmer) step4_a(){
	if z.ends(_AL) {
		z.step4_update()
	}
}
func (z *stemmer) step4_c(){
	if z.ends(_ANCE) || z.ends(_ENCE) {
		z.step4_update()
	}

}
func (z *stemmer) step4_e(){
	if z.ends(_ER) {
		z.step4_update()
	}
}
func (z *stemmer) step4_i(){
	if z.ends(_IC) {
		z.step4_update()
	}
}
func (z *stemmer) step4_l(){
	if z.ends(_ABLE) || z.ends(_IBLE) {
		z.step4_update()
	}
}
func (z *stemmer) step4_n(){
	if z.ends(_ANT) || z.ends(_EMENT) || z.ends(_MENT) || z.ends(_ENT) {
		z.step4_update()
	}
}
func (z *stemmer) step4_o(){
	if z.ends(_OU) {
		z.step4_update()
	}
	if z.ends(_ION) && ('s' == z.b[z.j] || 't' == z.b[z.j]) {
		z.step4_update()
	}
}
func (z *stemmer) step4_s(){
	if z.ends(_ISM) {
		z.step4_update()
	}
}
func (z *stemmer) step4_t(){
	if z.ends(_ATE) || z.ends(_ITI) {
		z.step4_update()
	}
}
func (z *stemmer) step4_u(){
	if z.ends(_OUS) {
		z.step4_update()
	}
}
func (z *stemmer) step4_v(){
	if z.ends(_IVE) {
		z.step4_update()
	}
}
func (z *stemmer) step4_z(){
	if z.ends(_IZE) {
		z.step4_update()
	}
}

/* step5(z) removes a final -e if m(z) > 1, and changes -ll to -l if
   m(z) > 1. */


func (z *stemmer) step5() {
	z.j = z.k
	if 'e' == z.b[z.k] {
		a:=z.m()
		if 1<a || 1==a && !z.cvc(z.k-1) {
			z.k--
		}
	}
	if 'l' == z.b[z.k] && z.doublec(z.k) && 1 < z.m() {
		z.k--
	}
}

///* In stem(z, b, k), b is a char pointer, and the string to be stemmed is
//   from b[0] to b[k] inclusive.  Possibly b[k+1] == '\0', but it is not
//   important. The stemmer adjusts the characters b[0] ... b[k] and returns
//   the new end-point of the string, k'. Stemming never increases word
//   length, so 0 <= k' <= k.
//*/
func (z *stemmer) stem (b []byte) (int) {

	z.b = b
  z.j = 0
	z.k = len(b)-1

	if z.k > 1{
    z.step1ab()
    z.step1c()
    z.step2()
    z.step3()
    z.step4()
    z.step5()
  }
	return z.k
}

func (z *stemmer) String () string {
  return fmt.Sprintf ("stemmer {b=%s j=%d k=%d}", string(z.b), z.j, z.k)
}


func Stem(word string)(string) {
	var z stemmer
	b := []byte(strings.ToLower(word))
	bn := z.stem(b)
	if bn+1 <= len(z.b) {
		return (string)(z.b[:bn+1])
	}
  log ("%s %s \n", word, z.String())
	return ""
}


