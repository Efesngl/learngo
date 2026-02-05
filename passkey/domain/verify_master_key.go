package domain
type IMasterKeyVerifier interface{
	Verify(masterPassword []byte) error
}

func Execute() {

}
