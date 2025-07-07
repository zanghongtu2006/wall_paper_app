import 'package:flutter/material.dart';

import '../data/sample_wallpapers.dart';
import 'preview_page.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('精选壁纸')),
      body: ListView.separated(
        padding: const EdgeInsets.all(12),
        itemCount: sampleWallpapers.length,
        separatorBuilder: (_, __) => const SizedBox(height: 12),
        itemBuilder: (context, index) {
          final wall = sampleWallpapers[index];
          return GestureDetector(
            onTap: () => Navigator.push(
              context,
              MaterialPageRoute(
                builder: (_) => PreviewPage(wallpaper: wall),
              ),
            ),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Stack(
                children: [
                  Image.network(
                    wall.cover,
                    fit: BoxFit.cover,
                    width: double.infinity,
                    height: 250,
                    loadingBuilder: (c, w, p) => p == null
                        ? w
                        : Container(
                            color: Colors.grey[200],
                            height: 250,
                            child: const Center(
                                child: CircularProgressIndicator()),
                          ),
                  ),
                  Positioned(
                    bottom: 10,
                    right: 10,
                    child: Container(
                      color: Colors.black54,
                      padding: const EdgeInsets.symmetric(
                          horizontal: 6, vertical: 2),
                      child: Text(
                        wall.title,
                        style: const TextStyle(color: Colors.white),
                      ),
                    ),
                  ),
                ],
              ),
            ),
          );
        },
      ),
    );
  }
}
