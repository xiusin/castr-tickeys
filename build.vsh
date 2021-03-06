#!/usr/local/bin/v run

import term

term.clear()

println(term.ok_message('开始打包应用'))

system('qtdeploy build desktop')

println(term.ok_message('压缩执行文件'))

// system('upx -9 deploy/darwin/castr-tickeys.app/Contents/MacOS/castr-tickeys')

mkdir('deploy/darwin/castr-tickeys.app/Contents/Resources/sounds') or {
	println('sounds: ${term.fail_message(err)}')
	return
}

cp_all('sounds', 'deploy/darwin/castr-tickeys.app/Contents/Resources/sounds', true) or {
	println('sounds: ${term.fail_message(err)}')
	return
}

cp('package.json', 'deploy/darwin/castr-tickeys.app/Contents/Macos/package.json') or {
	println('package.json: ${term.fail_message(err)}')
	return
}

cp('castr-tickeys.icns', 'deploy/darwin/castr-tickeys.app/Contents/Resources/castr-tickeys.icns') ?

println(term.ok_message('构建完成!'))

system('./deploy/darwin/castr-tickeys.app/Contents/MacOS/castr-tickeys')
