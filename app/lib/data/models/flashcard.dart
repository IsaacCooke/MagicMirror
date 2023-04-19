import 'package:flutter/cupertino.dart';

@immutable
class FlashcardWidget extends StatelessWidget {
  final int id;
  final String term;
  final String definition;

  const FlashcardWidget({
    Key? key,
    required this.id,
    required this.term,
    required this.definition,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Row (
        children: [
          Expanded(
            flex: 1,
            child: term != null
              ? Text(term)
              : const Text("No data"),
          ),
          Expanded (
            flex: 2,
            child: Column(
              children: [
                Text(id.toString()),
                Text(term),
                Text(definition),
              ],
            ),
          ),
        ],
      ),
    );
  }
}