package disk

func isStable(fs string) bool {
	return fs == "ext" ||
		fs == "ext2" ||
		fs == "ext3" ||
		fs == "ext4" ||
		fs == "reiserfs" ||
		fs == "ntfs" ||
		fs == "msdos" ||
		fs == "dos" ||
		fs == "vfat" ||
		fs == "xfs" ||
		fs == "hpfs" ||
		fs == "jfs" ||
		fs == "ufs" ||
		fs == "hfs" ||
		fs == "hfsplus"
}
