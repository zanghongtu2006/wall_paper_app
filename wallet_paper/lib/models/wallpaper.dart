import 'dart:math';

enum WallpaperType { single, group }

class Wallpaper {
  final String id;
  final String title;
  final WallpaperType type;
  final List<String> imageUrls;
  final String author;

  const Wallpaper({
    required this.id,
    required this.title,
    required this.type,
    required this.imageUrls,
    required this.author,
  });

  String get cover => imageUrls.first;
  String get randomImage => (imageUrls[Random().nextInt(imageUrls.length)]);
}