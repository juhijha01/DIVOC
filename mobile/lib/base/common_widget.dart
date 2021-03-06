import 'package:divoc/base/constants.dart';
import 'package:divoc/generated/l10n.dart';
import 'package:flutter/material.dart';

class LoadingOverlay extends StatelessWidget {
  const LoadingOverlay({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.grey.withOpacity(0.5),
      child: Center(
        child: CircularProgressIndicator(),
      ),
    );
  }
}

class DivocHeader extends StatelessWidget {
  final bool showHeaderMenu;
  final bool showHelpMenu;

  DivocHeader({this.showHeaderMenu = false, this.showHelpMenu = false});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: SafeArea(
        child: SizedBox(
          height: 50,
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Expanded(
                flex: 7,
                child: Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Image(
                    image: AssetImage(ImageAssetPath.DIVOC_LOGO),
                  ),
                ),
              ),
              Visibility(
                visible: showHelpMenu,
                child: Expanded(
                  flex: 1,
                  child: SizedBox(
                    width: 25,
                    height: 25,
                    child: Image(
                      image: AssetImage(ImageAssetPath.HEADER_HELP),
                    ),
                  ),
                ),
              ),
              Visibility(
                visible: showHeaderMenu,
                child: Expanded(
                  flex: 2,
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: SizedBox(
                      width: 40,
                      height: 40,
                      child: Image(
                        image: AssetImage(ImageAssetPath.HEADER_MENU),
                      ),
                    ),
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}

class DivocFooter extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var themeData = Theme.of(context);
    return Container(
      color: themeData.primaryColorLight,
      child: Row(
        mainAxisSize: MainAxisSize.max,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Text(
              DivocLocalizations.of(context).tAndC,
              style: themeData.textTheme.caption
                  .copyWith(color: themeData.primaryColor),
            ),
          ),
        ],
      ),
    );
  }
}
