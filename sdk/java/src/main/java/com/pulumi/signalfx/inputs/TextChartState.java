// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TextChartState extends com.pulumi.resources.ResourceArgs {

    public static final TextChartState Empty = new TextChartState();

    /**
     * Description of the text note.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the text note.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Markdown text to display.
     * 
     */
    @Import(name="markdown")
    private @Nullable Output<String> markdown;

    /**
     * @return Markdown text to display.
     * 
     */
    public Optional<Output<String>> markdown() {
        return Optional.ofNullable(this.markdown);
    }

    /**
     * Name of the text note.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the text note.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Tags associated with the resource
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags associated with the resource
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The URL of the chart.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL of the chart.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private TextChartState() {}

    private TextChartState(TextChartState $) {
        this.description = $.description;
        this.markdown = $.markdown;
        this.name = $.name;
        this.tags = $.tags;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TextChartState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TextChartState $;

        public Builder() {
            $ = new TextChartState();
        }

        public Builder(TextChartState defaults) {
            $ = new TextChartState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the text note.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the text note.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param markdown Markdown text to display.
         * 
         * @return builder
         * 
         */
        public Builder markdown(@Nullable Output<String> markdown) {
            $.markdown = markdown;
            return this;
        }

        /**
         * @param markdown Markdown text to display.
         * 
         * @return builder
         * 
         */
        public Builder markdown(String markdown) {
            return markdown(Output.of(markdown));
        }

        /**
         * @param name Name of the text note.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the text note.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param url The URL of the chart.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the chart.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public TextChartState build() {
            return $;
        }
    }

}
