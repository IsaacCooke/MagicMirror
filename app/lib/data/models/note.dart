import 'package:flutter/cupertino.dart';

@immutable
class NoteWidget extends StatelessWidget {
  final int id;
  final String content;

  const NoteWidget({
    Key? key,
    required this.id,
    required this.content,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Row(
        children: [
          Expanded(
            flex: 1,
            child: content != null ? Text(content) : const Text("No data"),
          ),
          Expanded(
            flex: 2,
            child: Column(
              children: [
                Text(id.toString()),
                Text(content),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
