// id, title, body

import 'package:flutter/material.dart';

class PostRow extends StatelessWidget {
  final String id;
  final String title;
  final String body;

  const PostRow({
    Key? key,
    required this.id,
    required this.title,
    required this.body,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Row (
        children: [
          Expanded(
            flex: 1,
            child: body != null
              ? Text(body)
              : const Text('No body'),
          ),
          Expanded(
            flex: 2,
            child: Column (
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(id),
                Text(title),
              ],
            ),
          ),
        ],
      ),
    );
  }
}