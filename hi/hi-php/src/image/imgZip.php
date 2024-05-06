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
     * @param string $src  源地址
     * @param string $dist 目标地址
     * @param int $percent 缩放率
     */
    static function base($src, $dist, $percent=1){
        $imgInfo = getimagesize($src);
        $pInfo = pathinfo($src);

        list($width, $height, $type, $attr) = $imgInfo;
        $ext = image_type_to_extension($type, false);
        $stdExt = '.' . (isset($pInfo['extension'])? $pInfo['extension'] : '');
        if(strpos(strtolower($dist), $stdExt) === false){
            $dist .= $stdExt;
        }
        $imgCreateFn = 'imagecreatefrom'.$ext;

        $image = $imgCreateFn($src);
        $nWid = $width * $percent;
        $nHgt = $height * $percent;

        $image_thump = imagecreatetruecolor($nWid,$nHgt);
        //print_r([$nWid, $nHgt, $width, $height]);
        //将原图复制带图片载体上面，并且按照一定比例压缩,极大的保持了清晰度
        imagecopyresampled($image_thump, $image,0,0,0,0, $nWid, $nHgt, $width, $height);

        $saveFn = 'image'.$ext;
        $saveFn($image, $dist);

        imagedestroy($image);
    }

    /**
     * 生成缩略图
     * @param string $srcFile 源文件路径
     * @param string $dstFile 目标文件路径
     * @param int $size
     * @param bool $is_square
     * @param int $quality
     * @return bool
     */
    static function create_thumbnails($srcFile, $dstFile, $size = 150, $is_square = false, $quality = 100){
        if(file_exists($srcFile)){
            //返回含有4个单元的数组，0-宽，1-高，2-图像类型，3-宽高的文本描述。
            list(,, $type) = getimagesize($srcFile);

            $ext = image_type_to_extension($type, false);
            $imgCreateFn = 'imagecreatefrom'.$ext;
            $im = $imgCreateFn($srcFile);

            //设置标记以在保存 PNG 图像时保存完整的 alpha 通道信息（与单一透明色相反）
            imagesavealpha($im, true);
            //
            $srcW = imagesx($im);
            $srcH = imagesy($im);
            $srcX = $srcY = 0;
            if($is_square == true){
                if($srcH >= $srcW){
                    $srcX = 0;
                    $srcY = floor(($srcH - $srcW) / 2);
                    $srcH = $srcW;
                }else {
                    $srcY = 0;
                    $srcX = floor(($srcW - $srcH) / 2);
                    $srcW = $srcH;
                }
                $fdstH = $fdstW = $size;
            } else {
                if ($srcW < $size && $srcH < $size) {
                    return false;
                }
                if ($srcH >= $srcW) {
                    $fdstH = $size;
                    $fdstW = $fdstH * $srcW / $srcH;
                } else {
                    $fdstW = $size;
                    $fdstH = $fdstW * $srcH / $srcW;
                }
            }
            $ni = imagecreatetruecolor($fdstW, $fdstH);
            //关闭 alpha 渲染并设置 alpha 标志
            imagealphablending($ni, false);
            imagesavealpha($ni, true);
            //重采样拷贝部分图像并调整大小
            imagecopyresampled($ni, $im, 0, 0, $srcX, $srcY, $fdstW, $fdstH, $srcW, $srcH);

            $imgCreateFn = 'image'.$ext;
            $imgCreateFn($ni, $dstFile, $quality);

            imagedestroy($im);
            imagedestroy($ni);
        }
        return false;
    }
}

// 数据错误
// 测试
//imgZip::base(__DIR__.'/small.jpg', __DIR__.'/small-output.jpg');
//imgZip::base(__DIR__.'/middle.jpg', __DIR__.'/middle-output.jpg');

//imgZip::base(__DIR__.'/big.jpg', __DIR__.'/big-output.jpg', 0.02);
imgZip::create_thumbnails(__DIR__.'/big.jpg', __DIR__.'/big-output-create_thumbnails.jpg');

//imgZip::base(__DIR__.'/big.jpg', __DIR__.'/big-output-44.jpg', 0.3);