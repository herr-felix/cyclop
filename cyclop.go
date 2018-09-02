package cyclop

// Cyclop controls everything...
type Cyclop struct {
	Search  SearchProvider
	Cache   CacheProvider
	Storage StorageProvider
}
