package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"github.com/tendermint/tendermint/crypto"
	"math/big"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSKSuite(t *testing.T) {
	suite.Run(t, new(SKSuite))
}

type SKSuite struct{ CommonSuite }

func (suite *SKSuite) TestString() {
	const prefix = "abc"
	suite.Require().Equal(prefix+"{-}", suite.sk.String(prefix))
}

func (suite *SKSuite) TestPubKey() {
	pk := suite.sk.PubKey()
	suite.True(suite.sk.PublicKey.Equal(&pk.PublicKey))
}

func (suite *SKSuite) TestBytes() {
	bz := suite.sk.Bytes()
	suite.Len(bz, 32)
	var sk *PrivKey
	suite.Nil(sk.Bytes())
}

func (suite *SKSuite) TestMarshal() {
	require := suite.Require()
	const size = 32

	var buffer = make([]byte, size)
	suite.sk.MarshalTo(buffer)

	var sk = new(PrivKey)
	err := sk.Unmarshal(buffer, secp256r1, size)
	require.NoError(err)
	require.True(sk.Equal(&suite.sk.PrivateKey))
}

func (suite *SKSuite) TestSign() {
	require := suite.Require()

	msg := crypto.CRandBytes(1000)
	sig, err := suite.sk.Sign(msg)
	require.NoError(err)
	sigCpy := make([]byte, len(sig))
	copy(sigCpy, sig)
	require.True(suite.pk.VerifySignature(msg, sigCpy))


	for i := range sig {
		sigCpy[i] ^= byte(i + 1)
		require.False(suite.pk.VerifySignature(msg, sigCpy))
	}







	r := new(big.Int).SetBytes(sig[:32])
	low_s := new(big.Int).SetBytes(sig[32:64])



	require.Equal(NormalizeS(low_s), low_s)



	high_s := new(big.Int).Mod(new(big.Int).Neg(low_s), elliptic.P256().Params().N)

	require.False(suite.pk.VerifySignature(msg, signatureRaw(r, high_s)))


	sigCpy = make([]byte, len(sig)+2)
	copy(sigCpy, sig)
	sigCpy[65] = byte('A')

	require.False(suite.pk.VerifySignature(msg, sigCpy))



	hash := sha256.Sum256([]byte(msg))
	require.True(ecdsa.Verify(&suite.pk.PublicKey, hash[:], r, high_s))


	msg[1] ^= byte(2)
	require.False(suite.pk.VerifySignature(msg, sig))
}
