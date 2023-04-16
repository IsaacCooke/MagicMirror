import 'package:flutter/cupertino.dart';

const lightColorScheme = CupertinoThemeData(
  barBackgroundColor: Color(0xFFFEFBFF),
  brightness: Brightness.light,
  primaryColor: Color(0xFF305DA8),
  primaryContrastingColor: Color(0xFFFFFFFF),
  scaffoldBackgroundColor: Color(0xFFFEFBFF),
  textTheme: CupertinoTextThemeData(
    textStyle: TextStyle(
      color: Color(0xFF1B1B1F),
    ),
  ),
);

const darkColorScheme = CupertinoThemeData(
  barBackgroundColor: Color(0xFF1B1B1F),
  brightness: Brightness.dark,
  primaryColor: Color(0xFFADC6FF),
  primaryContrastingColor: Color(0xFF002E69),
  scaffoldBackgroundColor: Color(0xFF1B1B1F),
  textTheme: CupertinoTextThemeData(
    textStyle: TextStyle(
      color: Color(0xFFE3E2E6),
    ),
  ),
);