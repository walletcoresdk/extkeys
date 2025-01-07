package extkeys

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type VectorsFile struct {
	Data    map[string][][6]string
	vectors []*Vector
}

type Vector struct {
	language, password, input, mnemonic, seed, xprv string
}

func TestNewMnemonic(t *testing.T) {
	m1 := NewMnemonic()
	if m1.salt != defaultSalt {
		t.Errorf("expected default salt, got: %q", m1.salt)
	}
}

func TestMnemonic_WordList(t *testing.T) {
	m := NewMnemonic()
	_, err := m.WordList(EnglishLanguage)
	if err != nil {
		t.Errorf("expected WordList to return WordList without errors, got: %s", err)
	}

	indexes := []Language{-1, Language(len(m.wordLists))}

	fmt.Println(indexes)
	for _, index := range indexes {
		_, err := m.WordList(index)
		if err == nil {
			t.Errorf("expected WordList to return an error with index %d", index)
		}
	}
}

// TestMnemonicPhrase
func TestMnemonicPhrase(t *testing.T) {

	mnemonic := NewMnemonic()

	// test strength validation
	strengths := []EntropyStrength{127, 129, 257}
	for _, s := range strengths {
		mw, err := mnemonic.MnemonicPhrase(s, EnglishLanguage)
		fmt.Println(mw)
		if err != ErrInvalidEntropyStrength {
			t.Errorf("Entropy strength '%d' should be invalid", s)
		}
	}

	// test mnemonic generation
	t.Log("Test mnemonic generation:")
	for _, language := range mnemonic.AvailableLanguages() {
		phrase, err := mnemonic.MnemonicPhrase(EntropyStrength128, language)
		t.Logf("Mnemonic (%s): %s", Languages[language], phrase)

		if err != nil {
			t.Errorf("Test failed: could not create seed: %s", err)
		}

		if !mnemonic.ValidMnemonic(phrase, language) {
			t.Error("Seed is not valid Mnenomic")
		}
	}

	// run against test vectors
	vectorsFile, err := LoadVectorsFile("mnemonic_vectors.json")
	if err != nil {
		t.Error(err)
	}

	t.Log("Test against pre-computed seed vectors:")
	stats := map[string]int{}
	for _, vector := range vectorsFile.vectors {
		stats[vector.language]++
		mnemonic := NewMnemonic()
		seed := mnemonic.MnemonicSeed(vector.mnemonic, vector.password)
		if fmt.Sprintf("%x", seed) != vector.seed {
			t.Errorf("Test failed (%s): incorrect seed (%x) generated (expected: %s)", vector.language, seed, vector.seed)
			return
		}
		//t.Logf("Test passed: correct seed (%x) generated (expected: %s)", seed, vector.seed)

		rootKey, err := NewMaster(seed)
		if err != nil {
			t.Error(err)
		}

		if rootKey.String() != vector.xprv {
			t.Errorf("Test failed (%s): incorrect xprv (%s) generated (expected: %s)", vector.language, rootKey, vector.xprv)
		}
	}
	for language, count := range stats {
		t.Logf("[%s]: %d tests completed", language, count)
	}
}

func LoadVectorsFile(path string) (*VectorsFile, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Test failed: cannot open vectors file: %s", err)
	}

	var vectorsFile VectorsFile
	if err := json.NewDecoder(fp).Decode(&vectorsFile); err != nil {
		return nil, fmt.Errorf("Test failed: cannot parse vectors file: %s", err)
	}

	// load data into Vector structs
	for language, data := range vectorsFile.Data {
		for _, item := range data {
			vectorsFile.vectors = append(vectorsFile.vectors, &Vector{language, item[0], item[1], item[2], item[3], item[4]})
		}
	}

	return &vectorsFile, nil
}

func (v *Vector) String() string {
	return fmt.Sprintf("{password: %s, input: %s, mnemonic: %s, seed: %s, xprv: %s}",
		v.password, v.input, v.mnemonic, v.seed, v.xprv)
}

// TestMnemonicPhrase
func TestMnemonicIndex(t *testing.T) {

	mnemonic := NewMnemonic()

	hrase, err := mnemonic.MnemonicPhrase(EntropyStrength128, EnglishLanguage)
	if err != nil {
		t.Logf("MnemonicPhrase failed: %s", err)
		return
	}
	t.Logf("Mnemonic : %s", hrase)

	raw, err := mnemonic.GetIndexesFromeWords(hrase, EnglishLanguage)
	if err != nil {
		t.Logf("MnemonicPhrase failed: %s", err)
		return
	}
	t.Logf("indexes : %v", raw)

	words1, err := mnemonic.RestoreWordsFromHex(raw.Hex, EnglishLanguage)
	if err != nil {
		t.Logf("MnemonicPhrase failed: %s", err)
		return
	}
	t.Logf("words : %v", words1)

	words2, err := mnemonic.RestoreWordsFromIndexes(raw.Indexes, EnglishLanguage)
	if err != nil {
		t.Logf("MnemonicPhrase failed: %s", err)
		return
	}
	t.Logf("words : %v", words2)

	// 给定 Entropy
	entropyHex := "01696082f6c0c606d44f12b8d5dad54f"
	entropy, err := hex.DecodeString(entropyHex)
	if err != nil {
		fmt.Println("Invalid entropy:", err)
		return
	}

	// 生成助记词
	mnemonicV2, err := mnemonic.MnemonicEntropy(entropy)
	if err != nil {
		fmt.Println("Error generating mnemonic:", err)
		return
	}

	fmt.Println("Generated mnemonicV2:", mnemonicV2)

	mnemonicV3 := "accident enrich camera unique arrest address eye time rhythm put relief oxygen"
	entropyV2, err := mnemonic.MnemonicToEntropy(mnemonicV3)
	if err != nil {
		fmt.Println("Error generating mnemonic:", err)
		return
	}

	fmt.Printf("Generated entropyV2: %x\n", entropyV2)
	fmt.Println("Generated entropyV2:", entropyV2)

}
