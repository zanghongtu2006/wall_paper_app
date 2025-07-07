import '../models/wallpaper.dart';

typedef WallpaperList = List<Wallpaper>;

const sampleWallpapers = [
  Wallpaper(
    id: '1',
    title: '清晨雪山',
    type: WallpaperType.single,
    imageUrls: ['https://picsum.photos/id/1015/800/1600'],
    author: '摄影师 A',
  ),
  Wallpaper(
    id: '2',
    title: '幻想都市组图',
    type: WallpaperType.group,
    imageUrls: [
      'https://picsum.photos/id/1018/800/1600',
      'https://picsum.photos/id/1019/800/1600',
      'https://picsum.photos/id/1020/800/1600',
    ],
    author: '设计师 B',
  ),
];
