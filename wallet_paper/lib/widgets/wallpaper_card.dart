import 'package:flutter/material.dart';

import '../models/wallpaper_card.dart';

class WallpaperCard extends StatelessWidget {
  final Wallpaper wallpaper;
  final VoidCallback onTap;

  const WallpaperCard({required this.wallpaper, required this.onTap});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: ClipRRect(
        borderRadius: BorderRadius.circular(12),
        child: Stack(
          children: [
            Image.network(
              wallpaper.imageUrl,
              fit: BoxFit.cover,
              width: double.infinity,
              height: 250,
              loadingBuilder: (context, child, progress) {
                if (progress == null) return child;
                return Container(
                  color: Colors.grey[200],
                  child: Center(child: CircularProgressIndicator()),
                );
              },
            ),
            Positioned(
              bottom: 10,
              right: 10,
              child: Container(
                color: Colors.black54,
                padding: EdgeInsets.symmetric(horizontal: 6, vertical: 2),
                child: Text(
                  wallpaper.author,
                  style: TextStyle(color: Colors.white),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
