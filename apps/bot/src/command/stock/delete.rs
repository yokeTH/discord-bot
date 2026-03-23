use ::serenity::all::{
    CreateActionRow, CreateSelectMenu, CreateSelectMenuKind, CreateSelectMenuOption,
};
use anyhow::bail;
use tracing::{debug, info, instrument};

use crate::component::SELECT_DELETE_ID;
use crate::{Context, Error};

#[poise::command(slash_command)]
#[instrument(name = "cmd_delete", skip(ctx), fields(user_id = %ctx.author().id))]
pub async fn delete(ctx: Context<'_>) -> Result<(), Error> {
    ctx.defer().await?;
    debug!("deferred reply");

    let symbol_store = ctx.data().symbol_store.clone();

    let symbols: Vec<String> = symbol_store.list().await?;
    if symbols.is_empty() {
        info!("attempted delete from empty watchlist");
        bail!("Watchlist is empty.");
    }

    let limit = symbols.len().min(25);

    let opts: Vec<CreateSelectMenuOption> = symbols
        .into_iter()
        .take(limit)
        .map(|sym: String| CreateSelectMenuOption::new(sym.clone(), sym))
        .collect();

    let menu = CreateSelectMenu::new(
        SELECT_DELETE_ID,
        CreateSelectMenuKind::String { options: opts },
    )
    .placeholder("Choose symbols...")
    .min_values(1)
    .max_values(limit as u8);

    let components = vec![CreateActionRow::SelectMenu(menu)];

    info!(limit, "presenting symbols for deletion");

    ctx.send(
        poise::CreateReply::default()
            .content("Select symbols to delete (you can pick multiple):")
            .components(components),
    )
    .await?;

    info!("sent selection menu");
    Ok(())
}
