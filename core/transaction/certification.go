package transaction

import (
	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	corepb "github.com/medibloc/go-medibloc/core/pb"
	coreState "github.com/medibloc/go-medibloc/core/state"
	"github.com/medibloc/go-medibloc/util/byteutils"
)

// AddCertificationPayload is payload type for AddCertificationTx
type AddCertificationPayload struct {
	IssueTime       int64
	ExpirationTime  int64
	CertificateHash []byte
}

// FromBytes converts bytes to payload.
func (payload *AddCertificationPayload) FromBytes(b []byte) error {
	payloadPb := &corepb.AddCertificationPayload{}
	if err := proto.Unmarshal(b, payloadPb); err != nil {
		return err
	}
	payload.IssueTime = payloadPb.IssueTime
	payload.ExpirationTime = payloadPb.ExpirationTime
	payload.CertificateHash = payloadPb.Hash
	return nil
}

// ToBytes returns marshaled AddCertificationPayload
func (payload *AddCertificationPayload) ToBytes() ([]byte, error) {
	payloadPb := &corepb.AddCertificationPayload{
		IssueTime:      payload.IssueTime,
		ExpirationTime: payload.ExpirationTime,
		Hash:           payload.CertificateHash,
	}
	return proto.Marshal(payloadPb)
}

//AddCertificationTx is a structure for adding certification
type AddCertificationTx struct {
	Issuer          common.Address
	Certified       common.Address
	CertificateHash []byte
	IssueTime       int64
	ExpirationTime  int64
	size            int
}

//NewAddCertificationTx returns AddCertificationTx
func NewAddCertificationTx(tx *coreState.Transaction) (*ExecutableTx, error) {
	if len(tx.Payload()) > MaxPayloadSize {
		return nil, ErrTooLargePayload
	}
	payload := new(AddCertificationPayload)
	if err := BytesToTransactionPayload(tx.Payload(), payload); err != nil {
		return nil, err
	}
	size, err := tx.Size()
	if err != nil {
		return nil, err
	}
	if !common.IsHexAddress(tx.From().Hex()) || !common.IsHexAddress(tx.To().Hex()) {
		return nil, ErrInvalidAddress
	}
	if !common.IsHash(byteutils.Bytes2Hex(payload.CertificateHash)) {
		return nil, ErrCertHashInvalid
	}

	return &ExecutableTx{
		Transaction: tx,
		Executable: &AddCertificationTx{
			Issuer:          tx.From(),
			Certified:       tx.To(),
			CertificateHash: payload.CertificateHash,
			IssueTime:       payload.IssueTime,
			ExpirationTime:  payload.ExpirationTime,
			size:            size,
		},
	}, nil
}

//Execute AddCertificationTx
func (tx *AddCertificationTx) Execute(bs blockState) error {
	certified, err := bs.GetAccount(tx.Certified)
	if err != nil {
		return err
	}
	_, err = certified.GetData(coreState.CertReceivedPrefix, tx.CertificateHash)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == nil {
		return ErrCertReceivedAlreadyAdded
	}

	issuer, err := bs.GetAccount(tx.Issuer)
	if err != nil {
		return err
	}
	_, err = issuer.GetData(coreState.CertIssuedPrefix, tx.CertificateHash)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == nil {
		return ErrCertIssuedAlreadyAdded
	}

	//TODO: certification payload Verify: drsleepytiger

	pbCertification := &corepb.Certification{
		CertificateHash: tx.CertificateHash,
		Issuer:          tx.Issuer.Bytes(),
		Certified:       tx.Certified.Bytes(),
		IssueTime:       tx.IssueTime,
		ExpirationTime:  tx.ExpirationTime,
		RevocationTime:  int64(-1),
	}
	certificationBytes, err := proto.Marshal(pbCertification)
	if err != nil {
		return err
	}

	// Add certification to certified's account state
	if err := certified.Data.Prepare(); err != nil {
		return err
	}
	if err := certified.Data.BeginBatch(); err != nil {
		return err
	}
	if err := certified.PutData(coreState.CertReceivedPrefix, tx.CertificateHash, certificationBytes); err != nil {
		if err := certified.Data.RollBack(); err != nil {
			return err
		}
		return err
	}
	if err := certified.Data.Commit(); err != nil {
		return err
	}
	if err := certified.Data.Flush(); err != nil {
		return err
	}
	if err := bs.PutAccount(certified); err != nil {
		return err
	}

	// Add certification to issuer's account state
	issuer, err = bs.GetAccount(tx.Issuer)
	if err != nil {
		return err
	}
	if err := issuer.Data.Prepare(); err != nil {
		return err
	}
	if err := issuer.Data.BeginBatch(); err != nil {
		return err
	}
	if err := issuer.PutData(coreState.CertIssuedPrefix, tx.CertificateHash, certificationBytes); err != nil {
		if err := issuer.Data.RollBack(); err != nil {
			return err
		}
		return err
	}
	if err := issuer.Data.Commit(); err != nil {
		return err
	}
	if err := issuer.Data.Flush(); err != nil {
		return err
	}
	if err := bs.PutAccount(issuer); err != nil {
		return err
	}

	return nil
}

//Bandwidth returns bandwidth.
func (tx *AddCertificationTx) Bandwidth() *common.Bandwidth {
	return common.NewBandwidth(1500, uint64(tx.size))
}

//RevokeCertificationTx is a structure for revoking certification
type RevokeCertificationTx struct {
	Revoker         common.Address
	CertificateHash []byte
	size            int
}

// RevokeCertificationPayload is payload type for RevokeCertificationTx
type RevokeCertificationPayload struct {
	CertificateHash []byte
}

// FromBytes converts bytes to payload.
func (payload *RevokeCertificationPayload) FromBytes(b []byte) error {
	payloadPb := &corepb.RevokeCertificationPayload{}
	if err := proto.Unmarshal(b, payloadPb); err != nil {
		return err
	}
	payload.CertificateHash = payloadPb.Hash
	return nil
}

// ToBytes returns marshaled RevokeCertificationPayload
func (payload *RevokeCertificationPayload) ToBytes() ([]byte, error) {
	payloadPb := &corepb.RevokeCertificationPayload{
		Hash: payload.CertificateHash,
	}
	return proto.Marshal(payloadPb)
}

//NewRevokeCertificationTx returns RevokeCertificationTx
func NewRevokeCertificationTx(tx *coreState.Transaction) (*ExecutableTx, error) {
	if len(tx.Payload()) > MaxPayloadSize {
		return nil, ErrTooLargePayload
	}
	payload := new(RevokeCertificationPayload)
	if err := BytesToTransactionPayload(tx.Payload(), payload); err != nil {
		return nil, err
	}
	size, err := tx.Size()
	if err != nil {
		return nil, err
	}
	if !common.IsHexAddress(tx.From().Hex()) {
		return nil, ErrInvalidAddress
	}
	if !common.IsHash(byteutils.Bytes2Hex(payload.CertificateHash)) {
		return nil, ErrCertHashInvalid
	}

	return &ExecutableTx{
		Transaction: tx,
		Executable: &RevokeCertificationTx{
			Revoker:         tx.From(),
			CertificateHash: payload.CertificateHash,
			size:            size,
		},
	}, nil
}

//Execute RevokeCertificationTx
func (tx *RevokeCertificationTx) Execute(bs blockState) error {
	issuer, err := bs.GetAccount(tx.Revoker)
	if err != nil {
		return err
	}
	certBytes, err := issuer.GetData(coreState.CertIssuedPrefix, tx.CertificateHash)
	if err != nil {
		return err
	}

	pbCert := new(corepb.Certification)
	err = proto.Unmarshal(certBytes, pbCert)
	if err != nil {
		return err
	}
	// verify transaction
	if !byteutils.Equal(pbCert.Issuer, tx.Revoker.Bytes()) {
		return ErrCertRevokerInvalid
	}
	if pbCert.RevocationTime > int64(-1) {
		return ErrCertAlreadyRevoked
	}
	if pbCert.ExpirationTime < bs.Timestamp() {
		return ErrCertAlreadyExpired
	}

	pbCert.RevocationTime = bs.Timestamp()
	newCertBytes, err := proto.Marshal(pbCert)
	if err != nil {
		return err
	}
	// change cert on issuer's cert issued List
	err = issuer.Data.Prepare()
	if err != nil {
		return err
	}
	err = issuer.Data.BeginBatch()
	if err != nil {
		return err
	}
	err = issuer.PutData(coreState.CertIssuedPrefix, tx.CertificateHash, newCertBytes)
	if err != nil {
		return err
	}
	err = issuer.Data.Commit()
	if err != nil {
		return err
	}
	err = issuer.Data.Flush()
	if err != nil {
		return err
	}
	err = bs.PutAccount(issuer)
	if err != nil {
		return err
	}
	// change cert on certified's cert received list
	certAddr, err := common.BytesToAddress(pbCert.Certified)
	if err != nil {
		return err
	}
	certified, err := bs.GetAccount(certAddr)
	if err != nil {
		return err
	}
	err = certified.Data.Prepare()
	if err != nil {
		return err
	}
	err = certified.Data.BeginBatch()
	if err != nil {
		return err
	}
	err = certified.PutData(coreState.CertReceivedPrefix, tx.CertificateHash, newCertBytes)
	if err != nil {
		return err
	}
	err = certified.Data.Commit()
	if err != nil {
		return err
	}
	err = certified.Data.Flush()
	if err != nil {
		return err
	}
	return bs.PutAccount(certified)
}

//Bandwidth returns bandwidth.
func (tx *RevokeCertificationTx) Bandwidth() *common.Bandwidth {
	return common.NewBandwidth(1500, uint64(tx.size))
}