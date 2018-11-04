package ffmpeg

// ffmpeg -y \
// -loop 1 -t 15 -i input1.jpg \
// -loop 1 -t 15 -i input2.jpg \
// -filter_complex \
// "[0:v]scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,setsar=1,fade=t=in:st=0:d=1,fade=t=out:st=14:d=1[v0]; \
//  [1:v]scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,setsar=1,fade=t=in:st=0:d=1,fade=t=out:st=14:d=1[v1]; \
//  [v0][v1]concat=n=2:v=1:a=0,format=yuv420p[v]" -map "[v]" out.mp4

func ImageLoopFadeToBlack(image1Path string, image2Path string, outputPath string, duration int) {
	command := "ffmpeg -y "
	command += "-loop 1 -t " + duration + " -i " + image1Path + " "
	command += "-loop 1 -t 15 -i " + image2Path + " "
	command += "-filter_complex "
	command += "\"[0:v]scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,setsar=1,fade=t=in:st=0:d=1,fade=t=out:st=14:d=1[v0]; "
	command += "[1:v]scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,setsar=1,fade=t=in:st=0:d=1,fade=t=out:st=14:d=1[v1]; "
	command += "[v0][v1]concat=n=2:v=1:a=0,format=yuv420p[v]\" -map \"[v]\" out.mp4"
}
