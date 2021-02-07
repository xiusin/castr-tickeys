#!/usr/local/bin/v run

import term

term.clear()

println ("转码各个音频")

walk("sounds", fn (wav string) {
    dirname := dir(wav)
    name := file_name(wav)
    outname := "output_" + dirname +"/" + name [0..name.len - 4] + ".mp3"
    system("ffmpeg -i ${wav} -f mp3 -acodec libmp3lame -y ${outname}")
})

