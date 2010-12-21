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
	_BLANK = getBytes("")
	ABLE = getBytes("able")
	AL  = getBytes("al")
	ALISM = getBytes("alism")
	ALITI = getBytes("aliti")
	ALIZE = getBytes("alize")
	ALLI = getBytes("alli")
	ANCE = getBytes("ance")
	ANCI = getBytes("anci")
	ANT = getBytes("ant")
	AT  = getBytes("at")
	ATE = getBytes("ate")
	ATION = getBytes("ation")
	ATIONAL = getBytes("ational")
	ATIVE = getBytes("ative")
	ATOR = getBytes("ator")
	BILITI = getBytes("biliti")
	BL  = getBytes("bl")
	BLE = getBytes("ble")
	BLI = getBytes("bli")
	E   = getBytes("e")
	ED   = getBytes("ed")
	EED  = getBytes("eed")
	ELI = getBytes("eli")
	EMENT = getBytes("ement")
	ENCE = getBytes("ence")
	ENCI = getBytes("enci")
	ENT = getBytes("ent")
	ENTLI = getBytes("entli")
	ER = getBytes("er")
	FUL= getBytes("ful")
	FULNESS = getBytes("fulness")
	I    = getBytes("i")
	IBLE = getBytes("ible")
	IC = getBytes("ic")
	ICAL = getBytes("ical")
	ICATE = getBytes("icate")
	ICITI = getBytes("iciti")
	IES  = getBytes("ies")
	ING  = getBytes("ing")
	ION = getBytes("ion")
	ISM = getBytes("ism")
	ITI = getBytes("iti")
	IVE = getBytes("ive")
	IVENESS = getBytes("iveness")
	IVITI = getBytes("iviti")
	IZ = getBytes("iz")
	IZATION = getBytes("ization")
	IZE = getBytes("ize")
	IZER = getBytes("izer")
	LOG = getBytes("log")
	LOGI = getBytes("logi")
	MENT = getBytes("ment")
	NESS = getBytes("ness")
	OU = getBytes("ou")
	OUS = getBytes("ous")
	OUSLI = getBytes("ousli")
	OUSNESS = getBytes("ousness")
	SSES = getBytes("sses")
	TION = getBytes("tion")
	TIONAL = getBytes("tional")
	Y    = getBytes("y")

 
)

func getBytes (s string)([]byte) {
	buf := bytes.NewBufferString(s)
	return buf.Bytes()
}

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
  //log("cvc -> %s %d\n", string(z.b), i)
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
	for i:=0; i!=z.j; i++ {
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
      //log("here %s %d\n", string(z.b), i) 
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
	//length := len(s)
	j      := z.j

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
			case z.ends(SSES):
				z.k -= 2
			case z.ends(IES):
				z.setto(I)
			default:
				if 's' != z.b[z.k-1] {
					z.k--
				}
		}
	}
	if z.ends(EED) {
		if 0 < z.m() {
			z.k--
		}
	} else if (z.ends(ED) || z.ends(ING)) && z.vowelinstem() {
		z.k = z.j
		switch {
			case z.ends(AT):
				z.setto(ATE)
			case z.ends(BL):
				z.setto(BLE)
			case z.ends(IZ):
				z.setto(IZE)
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
					z.setto(E)
				}
		}
	}
}

/* step1c(z) turns terminal y to i when there is another vowel in the stem. */


func (z *stemmer) step1c() {
	if z.ends(Y) && z.vowelinstem() {
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
		case z.ends(ATIONAL): z.r(ATE)
		case z.ends(TIONAL):  z.r(TION)
	}
}
func (z *stemmer) step2_c(){
	switch {
		case z.ends(ENCI): z.r(ENCE)
		case z.ends(ANCI): z.r(ANCE)
	}
}
func (z *stemmer) step2_e(){
	if z.ends(IZER) {
		z.r(IZE)
	}
}
func (z *stemmer) step2_l(){
	switch {
		case z.ends(BLI):   z.r(BLE)
		case z.ends(ALLI):  z.r(AL)
		case z.ends(ENTLI): z.r(ENT)
		case z.ends(ELI):   z.r(E)
		case z.ends(OUSLI): z.r(OUS)
	}
}
func (z *stemmer) step2_o(){
	switch {
		case z.ends(IZATION): z.r(IZE)
		case z.ends(ATION):   z.r(ATE)
		case z.ends(ATOR):    z.r(ATE)
	}
}

func (z *stemmer) step2_s(){
	switch {
		case z.ends(ALISM):   z.r(AL)
		case z.ends(IVENESS): z.r(IVE)
		case z.ends(FULNESS): z.r(FUL)
		case z.ends(OUSNESS): z.r(OUS)
	}
}
func (z *stemmer) step2_t(){
	switch {
		case z.ends(ALITI):  z.r(AL)
		case z.ends(IVITI):  z.r(IVE)
		case z.ends(BILITI): z.r(BLE)
	}
}
func (z *stemmer) step2_g(){
	if z.ends(LOGI) {
		z.r(LOG)
	}
}

/* step3(z) deals with -ic-, -full, -ness etc. similar strategy to step2. */


func (z *stemmer) step3 () {
if z.k >= len(z.b){
  return
}

	switch z.b[z.k] {
		case'e':z.step3_e()
		case'i':z.step3_i()
		case'l':z.step3_l()
		case's':z.step3_s()
	}
}

func (z *stemmer) step3_e(){
	switch {
		case z.ends(ICATE): z.r(IC)
		case z.ends(ATIVE): z.r(_BLANK)
		case z.ends(ALIZE): z.r(AL)
	}
}
func (z *stemmer) step3_i(){
	if z.ends(ICITI) {
		z.r(IC)
	}
}
func (z *stemmer) step3_l(){
	switch {
		case z.ends(ICAL): z.r(IC)
		case z.ends(FUL): z.r(_BLANK)
	}
}
func (z *stemmer) step3_s(){
	if z.ends(NESS) {
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
	if z.ends(AL) {
		z.step4_update()
	}
}
func (z *stemmer) step4_c(){
	if z.ends(ANCE) || z.ends(ENCE) {
		z.step4_update()
	}

}
func (z *stemmer) step4_e(){
	if z.ends(ER) {
		z.step4_update()
	}
}
func (z *stemmer) step4_i(){
	if z.ends(IC) {
		z.step4_update()
	}
}
func (z *stemmer) step4_l(){
	if z.ends(ABLE) || z.ends(IBLE) {
		z.step4_update()
	}
}
func (z *stemmer) step4_n(){
	if z.ends(ANT) || z.ends(EMENT) || z.ends(MENT) || z.ends(ENT) {
		z.step4_update()
	}
}
func (z *stemmer) step4_o(){
	if z.ends(OU) {
		z.step4_update()
	}
	if z.ends(ION) && ('s' == z.b[z.j] || 't' == z.b[z.j]) {
		z.step4_update()
	}
}
func (z *stemmer) step4_s(){
	if z.ends(ISM) {
		z.step4_update()
	}
}
func (z *stemmer) step4_t(){
	if z.ends(ATE) || z.ends(ITI) {
		z.step4_update()
	}
}
func (z *stemmer) step4_u(){
	if z.ends(OUS) {
		z.step4_update()
	}
}
func (z *stemmer) step4_v(){
	if z.ends(IVE) {
		z.step4_update()
	}
}
func (z *stemmer) step4_z(){
	if z.ends(IZE) {
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
	if len (b) <= 1{
		return len(b)-1
	}
	z.b = b
  z.j = 0
	z.k = len(b)-1

	z.step1ab()
	z.step1c()
	z.step2()
	z.step3()
	z.step4()
	z.step5()
	return z.k
}

func (z *stemmer) String () string {
  return fmt.Sprintf ("stemmer {b=%s j=%d k=%d}", string(z.b), z.j, z.k)
}


func Stem(word string)(string) {
	var z stemmer
	b := getBytes(strings.ToLower(word))
	bn := z.stem(b)
	if bn+1 <= len(z.b) {
		return (string)(z.b[:bn+1])
	}
	return ""
}

