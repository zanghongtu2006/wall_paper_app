import 'package:flutter/material.dart';

import '../models/wallpaper.dart';
import '../services/wallpaper_service.dart';
class PreviewPage extends StatelessWidget {
  final Wallpaper wallpaper;

  const PreviewPage({super.key, required this.wallpaper});

  @override
  Widget build(BuildContext context) {
    final url = wallpaper.type == WallpaperType.group ? wallpaper.randomImage : wallpaper.cover;

    return Scaffold(
      appBar: AppBar(title: Text(wallpaper.title)),
      body: Stack(
        fit: StackFit.expand,
        children: [
          Image.network(url, fit: BoxFit.cover),
          Positioned(
            bottom: 32,
            left: 16,
            right: 16,
            child: Column(
              children: [
                ElevatedButton.icon(
                  onPressed: () => WallpaperService.setHomeWallpaper(url),
                  icon: const Icon(Icons.wallpaper),
                  label: const Text('设置为桌面壁纸'),
                ),
                const SizedBox(height: 12),
                ElevatedButton.icon(
                  onPressed: () => WallpaperService.setLockScreen(url),
                  icon: const Icon(Icons.lock),
                  label: const Text('设置为锁屏壁纸'),
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}