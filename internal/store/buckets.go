package store

import bolt "go.etcd.io/bbolt"

var bucketNames = [][]byte{
	[]byte("registrations"), []byte("reviews"), []byte("records"),
	[]byte("audits"), []byte("workflows"), []byte("attachments"),
	[]byte("courses"), []byte("reviewers"), []byte("notifications"),
}

func ensureBuckets(tx *bolt.Tx) error {
	for _, name := range bucketNames {
		if _, err := tx.CreateBucketIfNotExists(name); err != nil {
			return err
		}
	}
	return nil
}
