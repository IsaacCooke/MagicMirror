import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/note.dart';

class GetNotes extends StatelessWidget {
  GetNotes({Key? key}) : super(key: key);

  static const String query = """
    query {
      getAllNotes {
        ID
        Content
      }
    }
  """;

  final Constants constants = Constants();

  Widget build(BuildContext context) {
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
                return NoteWidget(
                    id: id,
                    content: content
                );
              },
            );
          },
        ),
      ),
    );
  }
}