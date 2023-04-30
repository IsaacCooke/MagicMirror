import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/note.dart';

class DeleteNotes extends StatefulWidget {
  const DeleteNotes({Key? key}) : super(key: key);

  DeleteNotesState createState() => DeleteNotesState();
}

class DeleteNotesState extends State<DeleteNotes> {
  static const String query = """
    query {
      getAllNotes {
        ID
        Content
      }
    }
  """;

  void _deleteNote(int id) async {
    final client = constants.client;

    const String mutation = """
    mutation deleteNote(\$id: Int!) {
      deleteNote(id: \$id) {
        ID
      }
    }
    """;

    final options = MutationOptions(
      document: gql(mutation),
      variables: {
        "id": id,
      },
    );

    final result = await client.value.mutate(options);

    if(result.hasException) {
      throw Exception(result.exception.toString());
    }
  }

  final Constants constants = Constants();

  Widget build(BuildContext context){
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("Notes"),
      ),
      child: GraphQLProvider(
        client: constants.client,
        child: Query(
          options: QueryOptions(
            document: gql(query),
          ),
          builder: (result, {fetchMore, refetch}){
            if (result.isLoading){
              return const Center(
                child: CupertinoActivityIndicator(),
              );
            }
            if (result.data == null){
              return const Center(
                child: Text("No Data"),
              );
            }
            final notes = result.data!["getAllNotes"] as List<dynamic>;
            return ListView.builder(
              itemCount: notes.length,
              itemBuilder: (context, index){
                final note = notes[index];
                final int id = note["ID"];
                final String content = note["Content"];
                return Row(
                  mainAxisSize: MainAxisSize.min,
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Flexible(
                      child: NoteWidget(
                        id: id,
                        content: content,
                      ),
                    ),
                    CupertinoButton(
                      child: const Icon(CupertinoIcons.delete),
                      onPressed: () {
                        _deleteNote(
                          id,
                        );
                        refetch!();
                      },
                    ),
                  ],
                );
              },
            );
          },
        ),
      ),
    );
  }
}