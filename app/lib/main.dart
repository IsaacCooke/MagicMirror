import 'package:flutter/cupertino.dart';
import 'package:flutter/services.dart';

import 'package:app/pages/layout.dart';
import 'package:app/themes/cupertino_theme.dart';

void main() async {
  runApp(const App());
}

class App extends StatelessWidget {
  const App({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    SystemChrome.setPreferredOrientations([
      DeviceOrientation.portraitUp,
      DeviceOrientation.portraitDown,
    ]);
    return const CupertinoApp(
      debugShowCheckedModeBanner: false,
      theme: darkColorScheme,
      home: Layout(),
    );
  }
}