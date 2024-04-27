make run:
	(cd FtpClient; go run .; cd ../) &
	pushd ./FtpGUI/src/index.html;  python3 -m http.server 9999; popd;
	
