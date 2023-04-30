import 'package:flutter/cupertino.dart';

import 'package:app/pages/home.dart';
import 'package:app/pages/flashcards.dart';
import 'package:app/pages/notes.dart';
import 'package:app/pages/reminders.dart';

class Layout extends StatefulWidget {
  const Layout({Key? key}) : super(key: key);

  @override
  LayoutState createState() => LayoutState();
}

class LayoutState extends State<Layout> {

  @override
  Widget build(BuildContext context) {
    return CupertinoTabScaffold(
      tabBar: CupertinoTabBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.house_fill),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.list_bullet),
            label: 'Reminders',
          ),
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.doc_fill),
            label: 'Flashcards',
          ),
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.rectangle_paperclip),
            label: 'Notes',
          )
        ],
      ),
      tabBuilder: (context, index) {
        late final CupertinoTabView returnValue;
        switch (index) {
          case 0:
            returnValue = CupertinoTabView(builder: (context) {
              return const Home();
            });
            break;
          case 1:
            returnValue = CupertinoTabView(builder: (context) {
              return const Reminders();
            });
            break;
          case 2:
            returnValue = CupertinoTabView(builder: (context) {
              return const Flashcards();
            });
            break;
          case 3:
            returnValue = CupertinoTabView(builder: (context) {
              return const Notes();
            });
            break;
        }
        return returnValue;
      }
    );
  }
}