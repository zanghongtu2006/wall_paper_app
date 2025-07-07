import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'package:path_provider/path_provider.dart';
import 'package:wallpaper_manager_flutter/wallpaper_manager_flutter.dart';

class WallpaperService {
  static Future<void> setHomeWallpaper(String imageUrl) async {
    try {
      final file = await _downloadImage(imageUrl);
      await WallpaperManagerFlutter().setWallpaper(
        file,
        WallpaperManagerFlutter.homeScreen,
      );
    } catch (e) {
      // debugPrint('设置桌面壁纸失败: $e');
    }
  }

  static Future<void> setLockScreen(String imageUrl) async {
    try {
      final file = await _downloadImage(imageUrl);
      await WallpaperManagerFlutter().setWallpaper(
        file,
        WallpaperManagerFlutter.lockScreen,
      );
    } catch (e) {
      // debugPrint('设置锁屏壁纸失败: $e');
    }
  }

  static Future<File> _downloadImage(String url) async {
    final response = await http.get(Uri.parse(url));
    final tempDir = await getTemporaryDirectory();
    final file = File('${tempDir.path}/wall.jpg');
    await file.writeAsBytes(response.bodyBytes);
    return file;
  }
}
