<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/5/18
 * Time: 15:22
 * Email: conero@163.com
 * Name: 图片压缩实现
 */

class imgZip
{
    const imageTypeDick = [
        1 => 'GIF', 2 => 'JPG', 3 => 'PNG', 4 => 'SWF', 5 => 'PSD', 6 => 'BMP', 7 => 'TIFF', 8 => 'TIFF',
        9 => 'JPC', 10 => 'JP2', 11 => 'JPX', 12 => 'JB2', 13 => 'SWC', 14 => 'IFF', 15 => 'WBMP', 16 => 'XBM',
    ];

    /**
     * @param string $src
     * @param string $dist
     * @param int $percent
     */
    static function base($src, $dist, $percent=1){
        list($width, $height, $type, $attr) = getimagesize($src);
        $ext = image_type_to_extension($type, false);
        $stdExt = '.' . $ext;
        if(strpos(strtolower($dist), $stdExt) === false){
            $dist .= $stdExt;
        }
        $imgCreateFn = 'imagecreatefrom'.$ext;

        $image = $imgCreateFn($src);
        $nWid = $width * $percent;
        $nHgt = $height * $percent;

        $image_thump = imagecreatetruecolor($nWid,$nHgt);
        //将原图复制带图片载体上面，并且按照一定比例压缩,极大的保持了清晰度
        imagecopyresampled($image_thump, $image,0,0,0,0, $nWid, $nHgt, $width, $height);
        imagedestroy($image);

        $saveFn = 'image'.$ext;
        $saveFn($image, $dist);
    }
}

// 数据错误
// 测试
imgZip::base(__DIR__.'/small.jpg', __DIR__.'/small-output.jpg');
//imgZip::base(__DIR__.'/middle.jpg', __DIR__.'/middle-output.jpg');
//imgZip::base(__DIR__.'/big.jpg', __DIR__.'/big-output.jpg');